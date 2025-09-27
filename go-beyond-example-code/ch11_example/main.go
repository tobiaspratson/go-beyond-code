package main

import "fmt"

// Models
type Product struct {
    ID    int
    Name  string
    Price float64
}

type Order struct {
    ID       int
    UserID   int
    Products []Product
    Total    float64
}

// Interfaces
type ProductRepository interface {
    FindByID(id int) (*Product, error)
    FindAll() ([]Product, error)
}

type OrderRepository interface {
    Save(order *Order) error
    FindByUserID(userID int) ([]Order, error)
}

type PaymentService interface {
    ProcessPayment(amount float64) error
}

type EmailService interface {
    SendOrderConfirmation(order *Order) error
}

// Implementations
type InMemoryProductRepository struct {
    products map[int]*Product
}

func NewInMemoryProductRepository() ProductRepository {
    products := map[int]*Product{
        1: {ID: 1, Name: "Laptop", Price: 999.99},
        2: {ID: 2, Name: "Mouse", Price: 29.99},
        3: {ID: 3, Name: "Keyboard", Price: 79.99},
    }
    return &InMemoryProductRepository{products: products}
}

func (r *InMemoryProductRepository) FindByID(id int) (*Product, error) {
    product, exists := r.products[id]
    if !exists {
        return nil, fmt.Errorf("product not found")
    }
    return product, nil
}

func (r *InMemoryProductRepository) FindAll() ([]Product, error) {
    var products []Product
    for _, product := range r.products {
        products = append(products, *product)
    }
    return products, nil
}

type InMemoryOrderRepository struct {
    orders map[int]*Order
    nextID int
}

func NewInMemoryOrderRepository() OrderRepository {
    return &InMemoryOrderRepository{
        orders: make(map[int]*Order),
        nextID: 1,
    }
}

func (r *InMemoryOrderRepository) Save(order *Order) error {
    if order.ID == 0 {
        order.ID = r.nextID
        r.nextID++
    }
    r.orders[order.ID] = order
    return nil
}

func (r *InMemoryOrderRepository) FindByUserID(userID int) ([]Order, error) {
    var userOrders []Order
    for _, order := range r.orders {
        if order.UserID == userID {
            userOrders = append(userOrders, *order)
        }
    }
    return userOrders, nil
}

type MockPaymentService struct{}

func (m *MockPaymentService) ProcessPayment(amount float64) error {
    fmt.Printf("Processing payment of $%.2f\n", amount)
    return nil
}

type MockEmailService struct{}

func (m *MockEmailService) SendOrderConfirmation(order *Order) error {
    fmt.Printf("Sending order confirmation for order #%d\n", order.ID)
    return nil
}

// Order service with injected dependencies
type OrderService struct {
    productRepo ProductRepository
    orderRepo   OrderRepository
    paymentService PaymentService
    emailService  EmailService
}

func NewOrderService(
    productRepo ProductRepository,
    orderRepo OrderRepository,
    paymentService PaymentService,
    emailService EmailService,
) *OrderService {
    return &OrderService{
        productRepo:   productRepo,
        orderRepo:     orderRepo,
        paymentService: paymentService,
        emailService:  emailService,
    }
}

func (s *OrderService) CreateOrder(userID int, productIDs []int) (*Order, error) {
    var products []Product
    var total float64
    
    // Get products
    for _, id := range productIDs {
        product, err := s.productRepo.FindByID(id)
        if err != nil {
            return nil, fmt.Errorf("product %d not found: %v", id, err)
        }
        products = append(products, *product)
        total += product.Price
    }
    
    // Create order
    order := &Order{
        UserID:   userID,
        Products: products,
        Total:    total,
    }
    
    // Save order
    err := s.orderRepo.Save(order)
    if err != nil {
        return nil, err
    }
    
    // Process payment
    err = s.paymentService.ProcessPayment(total)
    if err != nil {
        return nil, err
    }
    
    // Send confirmation email
    err = s.emailService.SendOrderConfirmation(order)
    if err != nil {
        return nil, err
    }
    
    return order, nil
}

func main() {
    // Create dependencies
    productRepo := NewInMemoryProductRepository()
    orderRepo := NewInMemoryOrderRepository()
    paymentService := &MockPaymentService{}
    emailService := &MockEmailService{}
    
    // Create order service with injected dependencies
    orderService := NewOrderService(productRepo, orderRepo, paymentService, emailService)
    
    // Create an order
    order, err := orderService.CreateOrder(1, []int{1, 2})
    if err != nil {
        fmt.Printf("Error creating order: %v\n", err)
        return
    }
    
    fmt.Printf("Order created: %+v\n", order)
}