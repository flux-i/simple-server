package managers

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	et "simple-server/src/entities"
)

type RequestManager struct {
	initCache  map[int64]et.CompanyData   // Cache for initial data
	respCache  map[string]et.ResponseData // Cache for api data
	mutex      sync.Mutex
	processing map[string]bool // Track processing status of each request
}

func NewRequestManager() *RequestManager {
	return &RequestManager{
		initCache:  make(map[int64]et.CompanyData),
		respCache:  make(map[string]et.ResponseData),
		processing: make(map[string]bool),
	}
}

func (rm *RequestManager) ProcessRequest(req et.Request) {
	// Check if the request is already being processed
	key := fmt.Sprintf("%d_%s", req.CompanyID, req.API)
	rm.mutex.Lock()
	if rm.processing[key] {
		// Request is already being processed, wait for the response
		log.Println("Duplicate request with key:", key)

		for {
			_, ok := rm.respCache[key]
			if ok {
				break
			}
			// Sleep for a short duration before checking the cache again
			time.Sleep(time.Millisecond * 100)
		}
		data, _ := rm.respCache[key]
		// return the cache result
		req.Response <- data
		rm.mutex.Unlock()
		return
	}
	// Mark the request as being processed
	rm.processing[key] = true
	rm.mutex.Unlock()

	// Check if initial data is available in cache
	rm.mutex.Lock()
	initData, ok := rm.initCache[req.CompanyID]
	rm.mutex.Unlock()
	if !ok {
		log.Println("Getting initial data for CompanyID:", req.CompanyID, "API:", req.API)
		// Perform initial data calculation
		initData = CalculateInitialData(req.CompanyID)
		// Cache the initial data
		rm.mutex.Lock()
		rm.initCache[req.CompanyID] = initData
		rm.mutex.Unlock()
	}

	// Process the request with the initial data
	// Simulated calculation
	log.Println("Processing request for CompanyID:", req.CompanyID, "API:", req.API)
	var apiData interface{}
	switch req.API {
	case "financial":
		apiData = GenerateDummyFinancialData()
	case "sale":
		apiData = GenerateDummySalesData()
	case "employee":
		apiData = GenerateDummyEmployeesData()
	}
	data := et.ResponseData{InitData: initData, ApiData: apiData}
	// Simulated delay
	time.Sleep(time.Second * 2)
	// Cache the respone
	rm.respCache[key] = data
	// Send response
	req.Response <- data

	log.Println("Response sent")
	// Mark the request as processed
	rm.mutex.Lock()
	delete(rm.processing, key)
	delete(rm.respCache, key)
	rm.mutex.Unlock()
}

func CalculateInitialData(companyID int64) et.CompanyData {
	return et.CompanyData{
		CompanyID:  companyID,
		OtherValue: 1,
		ThirdValue: true,
	}
}

// GenerateDummyFinancialData generates dummy financial data for testing purposes.
func GenerateDummyFinancialData() et.FinancialData {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random values for financial metrics
	revenue := rand.Float64() * 1e6              // Random revenue between 0 and 1,000,000
	expenses := rand.Float64() * 500000          // Random expenses between 0 and 500,000
	profit := revenue - expenses                 // Calculate profit
	netIncome := profit * (rand.Float64() + 0.5) // Random net income between 50% to 150% of profit
	earningsPerShare := netIncome / 100          // Calculate earnings per share (for demonstration purposes)

	// Create and return FinancialData struct with dummy values
	return et.FinancialData{
		Revenue:          revenue,
		Expenses:         expenses,
		Profit:           profit,
		NetIncome:        netIncome,
		EarningsPerShare: earningsPerShare,
	}
}

// GenerateDummySalesData generates dummy sales data for testing purposes.
func GenerateDummySalesData() et.SalesData {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random values for sales metrics
	totalSales := rand.Float64() * 1e6              // Random total sales revenue between 0 and 1,000,000
	unitsSold := rand.Intn(1000) + 500              // Random number of units sold between 500 and 1500
	averagePrice := totalSales / float64(unitsSold) // Calculate average price per unit sold
	totalCustomers := rand.Intn(200) + 100          // Random number of customers between 100 and 300

	// Create and return SalesData struct with dummy values
	return et.SalesData{
		TotalSales:     totalSales,
		UnitsSold:      unitsSold,
		AveragePrice:   averagePrice,
		TotalCustomers: totalCustomers,
	}
}

// GenerateDummyEmployeeStats generates dummy employee statistics data for testing purposes.
func GenerateDummyEmployeesData() et.EmployeeStats {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random values for employee metrics
	totalEmployees := rand.Intn(500) + 100 // Random total number of employees between 100 and 600
	avgSalary := rand.Intn(5000) + 3000    // Random average salary between 3000 and 8000

	// Create and return EmployeeStats struct with dummy values
	return et.EmployeeStats{
		TotalEmployees: totalEmployees,
		AvgSalary:      avgSalary,
	}
}
