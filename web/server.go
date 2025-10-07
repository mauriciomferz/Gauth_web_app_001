package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Educational Demo Server for GAuth RFC-0150 Implementation
// ⚠️ EDUCATIONAL PURPOSE ONLY - NOT FOR PRODUCTION USE

type EducationalServer struct {
	router *gin.Engine
	port   string
}

type DemoResponse struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data,omitempty"`
	Educational bool        `json:"educational"`
	Timestamp   time.Time   `json:"timestamp"`
}

func NewEducationalServer(port string) *EducationalServer {
	// Set Gin to release mode for cleaner output
	gin.SetMode(gin.ReleaseMode)
	
	router := gin.New()
	
	// Add educational middleware
	router.Use(educationalMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	
	server := &EducationalServer{
		router: router,
		port:   port,
	}
	
	server.setupRoutes()
	return server
}

func educationalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add educational headers
		c.Header("X-Educational-Demo", "true")
		c.Header("X-GAuth-Version", "RFC-0150-Educational")
		c.Header("X-Warning", "Educational implementation - not for production use")
		
		// Add CORS headers for local development
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	}
}

func (s *EducationalServer) setupRoutes() {
	// Static files
	s.router.Static("/static", "./web/static")
	s.router.LoadHTMLGlob("./web/templates/*")
	
	// Main educational interface
	s.router.GET("/", s.serveIndex)
	
	// Educational API endpoints (simulated)
	api := s.router.Group("/api/v1/educational")
	{
		api.GET("/health", s.healthCheck)
		api.POST("/demo/token/create", s.demoCreateToken)
		api.POST("/demo/token/validate", s.demoValidateToken)
		api.POST("/demo/token/revoke", s.demoRevokeToken)
		api.POST("/demo/authz/check", s.demoAuthzCheck)
		api.GET("/demo/examples", s.listExamples)
		api.GET("/demo/architecture", s.getArchitecture)
	}
	
	// Documentation endpoints
	docs := s.router.Group("/docs")
	{
		docs.GET("/", s.serveDocs)
		docs.GET("/rfc", s.serveRFCInfo)
	}
}

func (s *EducationalServer) serveIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":       "GAuth Educational Demo",
		"Version":     "RFC-0150 Educational Implementation",
		"Timestamp":   time.Now().Format("2006-01-02 15:04:05"),
		"Educational": true,
	})
}

func (s *EducationalServer) healthCheck(c *gin.Context) {
	response := DemoResponse{
		Success:     true,
		Message:     "Educational demo server is running",
		Educational: true,
		Timestamp:   time.Now(),
		Data: map[string]interface{}{
			"version":     "RFC-0150-Educational",
			"environment": "learning",
			"uptime":      time.Since(time.Now()).String(),
			"warning":     "This is an educational implementation only",
		},
	}
	c.JSON(http.StatusOK, response)
}

func (s *EducationalServer) demoCreateToken(c *gin.Context) {
	// Simulate token creation for educational purposes
	time.Sleep(time.Millisecond * 500) // Simulate processing time
	
	token := map[string]interface{}{
		"id":        fmt.Sprintf("edu_token_%d", time.Now().Unix()),
		"type":      "educational_demo",
		"issuer":    "gauth-educational-demo",
		"subject":   "demo-user@example.com",
		"audience":  "learning-environment",
		"expiresAt": time.Now().Add(time.Hour).Unix(),
		"createdAt": time.Now().Unix(),
		"claims": map[string]interface{}{
			"scope":       "read write demo",
			"educational": true,
			"purpose":     "RFC-0150 demonstration",
		},
		"warning": "Educational token - not cryptographically secure",
	}
	
	response := DemoResponse{
		Success:     true,
		Message:     "Educational token created successfully",
		Data:        token,
		Educational: true,
		Timestamp:   time.Now(),
	}
	
	c.JSON(http.StatusOK, response)
}

func (s *EducationalServer) demoValidateToken(c *gin.Context) {
	// Simulate token validation
	time.Sleep(time.Millisecond * 300)
	
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, DemoResponse{
			Success:     false,
			Message:     "Invalid request format",
			Educational: true,
			Timestamp:   time.Now(),
		})
		return
	}
	
	// Simulate validation logic
	tokenId, exists := request["token_id"].(string)
	if !exists || tokenId == "" {
		c.JSON(http.StatusBadRequest, DemoResponse{
			Success:     false,
			Message:     "Token ID required for validation",
			Educational: true,
			Timestamp:   time.Now(),
		})
		return
	}
	
	validation := map[string]interface{}{
		"valid":      true,
		"token_id":   tokenId,
		"expires_at": time.Now().Add(time.Hour).Unix(),
		"claims_verified": []string{"scope", "educational", "purpose"},
		"warning":    "Educational validation - not production-grade security",
	}
	
	response := DemoResponse{
		Success:     true,
		Message:     "Token validation completed",
		Data:        validation,
		Educational: true,
		Timestamp:   time.Now(),
	}
	
	c.JSON(http.StatusOK, response)
}

func (s *EducationalServer) demoRevokeToken(c *gin.Context) {
	// Simulate token revocation
	time.Sleep(time.Millisecond * 400)
	
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, DemoResponse{
			Success:     false,
			Message:     "Invalid request format",
			Educational: true,
			Timestamp:   time.Now(),
		})
		return
	}
	
	tokenId, exists := request["token_id"].(string)
	if !exists || tokenId == "" {
		c.JSON(http.StatusBadRequest, DemoResponse{
			Success:     false,
			Message:     "Token ID required for revocation",
			Educational: true,
			Timestamp:   time.Now(),
		})
		return
	}
	
	revocation := map[string]interface{}{
		"revoked":           true,
		"token_id":          tokenId,
		"revocation_time":   time.Now().Unix(),
		"blacklist_added":   true,
		"sessions_invalidated": 1,
		"warning":           "Educational revocation - not persistent across restarts",
	}
	
	response := DemoResponse{
		Success:     true,
		Message:     "Token revoked successfully",
		Data:        revocation,
		Educational: true,
		Timestamp:   time.Now(),
	}
	
	c.JSON(http.StatusOK, response)
}

func (s *EducationalServer) demoAuthzCheck(c *gin.Context) {
	// Simulate authorization check
	time.Sleep(time.Millisecond * 350)
	
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, DemoResponse{
			Success:     false,
			Message:     "Invalid request format",
			Educational: true,
			Timestamp:   time.Now(),
		})
		return
	}
	
	action, _ := request["action"].(string)
	resource, _ := request["resource"].(string)
	
	// Simulate authorization decision
	allowed := action == "read" || action == "demo"
	policy := "educational_demo_policy"
	
	if action == "admin" {
		allowed = false
		policy = "deny_admin_in_demo"
	}
	
	authz := map[string]interface{}{
		"allowed":          allowed,
		"action":           action,
		"resource":         resource,
		"policy":           policy,
		"delegation_chain": []string{"user", "demo-session"},
		"evaluation_time":  time.Now().Unix(),
		"warning":          "Educational authorization - simplified logic for demonstration",
	}
	
	response := DemoResponse{
		Success:     true,
		Message:     "Authorization check completed",
		Data:        authz,
		Educational: true,
		Timestamp:   time.Now(),
	}
	
	c.JSON(http.StatusOK, response)
}

func (s *EducationalServer) listExamples(c *gin.Context) {
	examples := map[string]interface{}{
		"total_examples": 37,
		"categories": map[string]interface{}{
			"core": []string{
				"rfc_0115_poa_definition",
				"typed_events",
				"token_management",
				"authorization_engine",
			},
			"patterns": []string{
				"cascade_patterns",
				"resilience_patterns",
				"event_driven_architecture",
				"microservices_auth",
			},
			"advanced": []string{
				"distributed_systems",
				"legal_framework",
				"professional_interfaces",
				"audit_compliance",
			},
		},
		"status": "All examples are educational implementations",
		"repository": "https://github.com/Gimel-Foundation/GiFo-RFC-0150-Go-Implementation-of-GAuth-1.0/tree/main/examples",
	}
	
	response := DemoResponse{
		Success:     true,
		Message:     "Examples catalog retrieved",
		Data:        examples,
		Educational: true,
		Timestamp:   time.Now(),
	}
	
	c.JSON(http.StatusOK, response)
}

func (s *EducationalServer) getArchitecture(c *gin.Context) {
	architecture := map[string]interface{}{
		"layers": []string{
			"Web Interface (Educational)",
			"GAuth Core Engine",
			"Token Management",
			"Authorization Engine",
			"Event System",
			"Audit Trail",
			"Storage Backend",
		},
		"components": map[string]interface{}{
			"token_service": "JWT/PASETO token lifecycle management",
			"authz_engine":  "RBAC/ABAC with power-of-attorney flows",
			"event_bus":     "Typed event system with pub/sub patterns",
			"audit_logger":  "Structured audit trail for compliance",
			"config_mgr":    "Environment-specific configuration management",
		},
		"standards_compliance": []string{
			"GiFo-RFC-0111 (Power of Attorney)",
			"GiFo-RFC-0115 (Authorization Framework)", 
			"GiFo-RFC-0150 (Implementation Guidelines)",
		},
		"educational_notice": "This architecture represents learning concepts, not production deployment",
	}
	
	response := DemoResponse{
		Success:     true,
		Message:     "Architecture information retrieved",
		Data:        architecture,
		Educational: true,
		Timestamp:   time.Now(),
	}
	
	c.JSON(http.StatusOK, response)
}

func (s *EducationalServer) serveDocs(c *gin.Context) {
	docs := map[string]interface{}{
		"title": "GAuth Educational Documentation",
		"sections": []string{
			"Getting Started",
			"API Reference", 
			"Architecture Guide",
			"Examples Repository",
			"RFC Standards",
			"Compliance Implementation",
		},
		"disclaimer": "Educational documentation for learning purposes only",
		"links": map[string]string{
			"github":      "https://github.com/Gimel-Foundation/GiFo-RFC-0150-Go-Implementation-of-GAuth-1.0",
			"foundation":  "https://gimelfoundation.com",
			"rfc_repo":    "https://github.com/Gimel-Foundation/RFCs",
		},
	}
	
	c.JSON(http.StatusOK, docs)
}

func (s *EducationalServer) serveRFCInfo(c *gin.Context) {
	rfcInfo := map[string]interface{}{
		"implemented_rfcs": []map[string]string{
			{
				"id":          "GiFo-RFC-0111",
				"title":       "Power of Attorney Framework",
				"status":      "Educational Implementation",
				"description": "Defines power-of-attorney patterns for AI delegation",
			},
			{
				"id":          "GiFo-RFC-0115", 
				"title":       "Authorization Implementation",
				"status":      "Educational Implementation",
				"description": "Authorization engine with RBAC/ABAC support",
			},
			{
				"id":          "GiFo-RFC-0150",
				"title":       "Go Implementation Guidelines",
				"status":      "Educational Implementation", 
				"description": "Implementation patterns and best practices in Go",
			},
		},
		"compliance_level": "Educational demonstration of RFC concepts",
		"production_note":  "These implementations are for learning and should not be used in production environments",
	}
	
	c.JSON(http.StatusOK, rfcInfo)
}

func (s *EducationalServer) Start() error {
	fmt.Printf("\n🎓 GAuth Educational Demo Server\n")
	fmt.Printf("⚠️  EDUCATIONAL PURPOSE ONLY - NOT FOR PRODUCTION USE\n")
	fmt.Printf("📚 RFC-0150 Go Implementation Learning Environment\n\n")
	fmt.Printf("🌐 Server starting on: http://localhost%s\n", s.port)
	fmt.Printf("📖 Documentation: http://localhost%s/docs/\n", s.port)
	fmt.Printf("🔧 Health Check: http://localhost%s/api/v1/educational/health\n", s.port)
	fmt.Printf("\nPress Ctrl+C to stop the educational demo server\n\n")
	
	return s.router.Run(s.port)
}

func main() {
	// Educational demo server configuration
	port := ":8080"
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}
	
	// Create and start educational server
	server := NewEducationalServer(port)
	
	// Add educational startup message
	log.Printf("🎓 Starting GAuth Educational Demo Server")
	log.Printf("⚠️ Educational Implementation - Not for Production Use")
	
	if err := server.Start(); err != nil {
		log.Fatalf("❌ Failed to start educational demo server: %v", err)
	}
}