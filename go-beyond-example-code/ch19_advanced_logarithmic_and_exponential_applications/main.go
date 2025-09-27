package main

import (
    "fmt"
    "math"
)

// Calculate compound interest
func CompoundInterest(principal, rate, time float64) float64 {
    return principal * math.Exp(rate * time)
}

// Calculate compound interest with discrete compounding
func CompoundInterestDiscrete(principal, rate, time, compoundsPerYear float64) float64 {
    return principal * math.Pow(1 + rate/compoundsPerYear, compoundsPerYear * time)
}

// Calculate half-life decay
func HalfLifeDecay(initial, halfLife, time float64) float64 {
    return initial * math.Exp(-math.Ln2 * time / halfLife)
}

// Calculate pH from hydrogen ion concentration
func CalculatePH(hydrogenIonConcentration float64) float64 {
    return -math.Log10(hydrogenIonConcentration)
}

// Calculate hydrogen ion concentration from pH
func CalculateHydrogenIonConcentration(pH float64) float64 {
    return math.Pow(10, -pH)
}

// Calculate decibels (logarithmic scale for sound)
func CalculateDecibels(power, referencePower float64) float64 {
    return 10 * math.Log10(power / referencePower)
}

// Calculate magnitude of earthquake (Richter scale)
func CalculateRichterScale(amplitude, referenceAmplitude float64) float64 {
    return math.Log10(amplitude / referenceAmplitude)
}

// Calculate information entropy (Shannon entropy)
func CalculateEntropy(probabilities []float64) float64 {
    entropy := 0.0
    for _, p := range probabilities {
        if p > 0 {
            entropy -= p * math.Log2(p)
        }
    }
    return entropy
}

// Calculate logistic growth
func LogisticGrowth(initial, carryingCapacity, growthRate, time float64) float64 {
    return carryingCapacity / (1 + (carryingCapacity/initial - 1) * math.Exp(-growthRate * time))
}

// Calculate sigmoid function
func Sigmoid(x float64) float64 {
    return 1 / (1 + math.Exp(-x))
}

// Calculate hyperbolic functions using exponentials
func Sinh(x float64) float64 {
    return (math.Exp(x) - math.Exp(-x)) / 2
}

func Cosh(x float64) float64 {
    return (math.Exp(x) + math.Exp(-x)) / 2
}

func Tanh(x float64) float64 {
    return Sinh(x) / Cosh(x)
}

func main() {
    fmt.Println("=== Advanced Logarithmic and Exponential Applications ===")
    
    // Financial calculations
    fmt.Println("\n--- Financial Applications ---")
    principal := 1000.0
    rate := 0.05  // 5% annual rate
    time := 10.0  // 10 years
    
    // Continuous compounding
    continuous := CompoundInterest(principal, rate, time)
    fmt.Printf("Continuous compounding: $%.2f -> $%.2f\n", principal, continuous)
    
    // Discrete compounding
    annual := CompoundInterestDiscrete(principal, rate, time, 1)      // Annual
    monthly := CompoundInterestDiscrete(principal, rate, time, 12)     // Monthly
    daily := CompoundInterestDiscrete(principal, rate, time, 365)      // Daily
    
    fmt.Printf("Annual compounding: $%.2f -> $%.2f\n", principal, annual)
    fmt.Printf("Monthly compounding: $%.2f -> $%.2f\n", principal, monthly)
    fmt.Printf("Daily compounding: $%.2f -> $%.2f\n", principal, daily)
    
    // Scientific applications
    fmt.Println("\n--- Scientific Applications ---")
    
    // Half-life decay
    initial := 100.0
    halfLife := 5.0  // 5 years
    time = 10.0      // 10 years
    
    remaining := HalfLifeDecay(initial, halfLife, time)
    fmt.Printf("Half-life decay: %.2f -> %.2f (after %.1f years)\n", initial, remaining, time)
    
    // pH calculations
    hConcentration := 0.001  // 0.001 M
    ph := CalculatePH(hConcentration)
    fmt.Printf("pH of solution: %.2f\n", ph)
    
    // Reverse pH calculation
    calculatedH := CalculateHydrogenIonConcentration(ph)
    fmt.Printf("Hydrogen ion concentration from pH: %.6f M\n", calculatedH)
    
    // Sound and earthquake measurements
    fmt.Println("\n--- Logarithmic Scales ---")
    
    // Decibel calculation
    power := 1.0
    referencePower := 1e-12  // 1 picowatt
    db := CalculateDecibels(power, referencePower)
    fmt.Printf("Sound level: %.2f dB\n", db)
    
    // Richter scale
    amplitude := 1000.0
    referenceAmplitude := 1.0
    magnitude := CalculateRichterScale(amplitude, referenceAmplitude)
    fmt.Printf("Earthquake magnitude: %.2f\n", magnitude)
    
    // Information theory
    fmt.Println("\n--- Information Theory ---")
    probabilities := []float64{0.5, 0.25, 0.125, 0.125}
    entropy := CalculateEntropy(probabilities)
    fmt.Printf("Shannon entropy: %.4f bits\n", entropy)
    
    // Population growth models
    fmt.Println("\n--- Population Growth Models ---")
    initialPop := 100.0
    carryingCapacity := 1000.0
    growthRate := 0.1
    time = 20.0
    
    logisticPop := LogisticGrowth(initialPop, carryingCapacity, growthRate, time)
    fmt.Printf("Logistic growth: %.2f -> %.2f (after %.1f years)\n", initialPop, logisticPop, time)
    
    // Machine learning functions
    fmt.Println("\n--- Machine Learning Functions ---")
    x := 2.0
    sigmoid := Sigmoid(x)
    fmt.Printf("Sigmoid(%.1f): %.6f\n", x, sigmoid)
    
    // Hyperbolic functions using exponentials
    fmt.Println("\n--- Hyperbolic Functions (using exponentials) ---")
    testX := 1.0
    fmt.Printf("Sinh(%.1f): %.6f\n", testX, Sinh(testX))
    fmt.Printf("Cosh(%.1f): %.6f\n", testX, Cosh(testX))
    fmt.Printf("Tanh(%.1f): %.6f\n", testX, Tanh(testX))
    
    // Verify hyperbolic identity
    sinh := Sinh(testX)
    cosh := Cosh(testX)
    fmt.Printf("cosh²(x) - sinh²(x) = %.6f (should be 1)\n", cosh*cosh - sinh*sinh)
}