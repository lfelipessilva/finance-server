# 💰 Finance Server - Personal Finance Management API

[![Go Version](https://img.shields.io/badge/Go-1.23.6+-blue.svg)](https://golang.org)
[![Gin Framework](https://img.shields.io/badge/Gin-1.10.0+-green.svg)](https://gin-gonic.com)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Supported-blue.svg)](https://www.postgresql.org)
[![JWT Auth](https://img.shields.io/badge/JWT-Authentication-orange.svg)](https://jwt.io)
[![Google OAuth](https://img.shields.io/badge/Google-OAuth-red.svg)](https://developers.google.com/identity)

> **A powerful, scalable REST API for personal finance management built with Go, Gin, and PostgreSQL**

## 🚀 Features

### 💳 **Expense Management**

- ✅ **Create & Track Expenses** - Add expenses with categories, tags, and detailed information
- ✅ **Batch Operations** - Create, update, and delete multiple expenses at once
- ✅ **Advanced Filtering** - Filter by date, category, amount, and custom criteria
- ✅ **Smart Analytics** - Group expenses by day, month, year, and category
- ✅ **Multi-payment Support** - Track expenses across different banks and cards

### 📊 **Financial Analytics**

- 📈 **Spending Trends** - Analyze spending patterns over time
- 🏷️ **Category Analysis** - See where your money goes by category
- 📅 **Time-based Reports** - Daily, monthly, and yearly expense summaries
- 💰 **Total Calculations** - Automatic sum calculations with pagination

### 🔐 **Secure Authentication**

- 🔑 **JWT Token Authentication** - Secure API access with JSON Web Tokens
- 🌐 **Google OAuth Integration** - One-click login with Google accounts
- 👤 **User Management** - Create and manage user profiles
- 🛡️ **Protected Routes** - Sensitive operations require authentication

### 🏷️ **Organization Features**

- 📂 **Categories** - Organize expenses with customizable categories
- 🏷️ **Tags** - Add multiple tags to expenses for better organization
- 🔍 **Search & Filter** - Find expenses quickly with advanced search

### 🛠️ **Developer Friendly**

- 📚 **RESTful API** - Clean, intuitive REST endpoints
- 📖 **Comprehensive Documentation** - Complete API documentation with Bruno
- 🔧 **Easy Setup** - Simple configuration with environment variables
- 🐳 **Docker Ready** - Containerized deployment with Docker Compose

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   HTTP Layer    │    │  Business Logic │    │   Data Layer    │
│   (Gin Router)  │◄──►│   (Use Cases)   │◄──►│  (PostgreSQL)   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Handlers      │    │   Repositories  │    │   Migrations    │
│   (HTTP)        │    │   (Data Access) │    │   (Schema)      │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🚀 Quick Start

### Prerequisites

- Go 1.23.6+
- PostgreSQL 12+
- Docker (optional)

### 1. Clone the Repository

```bash
git clone https://github.com/lfelipessilva/finance-server.git
cd finance-server
```

### 2. Environment Setup

```bash
# Copy environment template
cp .env.example .env

# Edit with your database credentials
nano .env
```

### 3. Database Setup

```bash
# Run database migrations
./migrate.sh up
```

### 4. Start the Server

```bash
# Run the application
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080`

## 📚 API Endpoints

### 🔐 Authentication

| Method | Endpoint       | Description                    |
| ------ | -------------- | ------------------------------ |
| `POST` | `/api/v1/auth` | Authenticate with Google OAuth |

### 👤 User Management

| Method | Endpoint       | Description      |
| ------ | -------------- | ---------------- |
| `GET`  | `/api/v1/user` | Get user profile |
| `POST` | `/api/v1/user` | Create new user  |

### 💰 Expense Management

| Method   | Endpoint                 | Description              |
| -------- | ------------------------ | ------------------------ |
| `GET`    | `/api/v1/expenses`       | Get all expenses         |
| `POST`   | `/api/v1/expenses`       | Create single expense    |
| `POST`   | `/api/v1/expenses/batch` | Create multiple expenses |
| `PUT`    | `/api/v1/expenses/:id`   | Update expense           |
| `PUT`    | `/api/v1/expenses/batch` | Update multiple expenses |
| `DELETE` | `/api/v1/expenses/:id`   | Delete expense           |
| `DELETE` | `/api/v1/expenses/batch` | Delete multiple expenses |

### 📊 Analytics Endpoints

| Method | Endpoint                    | Description                |
| ------ | --------------------------- | -------------------------- |
| `GET`  | `/api/v1/expenses/category` | Group expenses by category |
| `GET`  | `/api/v1/expenses/date`     | Group expenses by date     |
| `GET`  | `/api/v1/expenses/day`      | Group expenses by day      |
| `GET`  | `/api/v1/expenses/month`    | Group expenses by month    |
| `GET`  | `/api/v1/expenses/year`     | Group expenses by year     |

### 📂 Categories & Tags

| Method | Endpoint             | Description        |
| ------ | -------------------- | ------------------ |
| `GET`  | `/api/v1/categories` | Get all categories |
| `GET`  | `/api/v1/tags`       | Get all tags       |

## 💡 Usage Examples

### Create an Expense

```bash
curl -X POST http://localhost:8080/api/v1/expenses \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Grocery Shopping",
    "description": "Weekly groceries from Walmart",
    "value": 85.50,
    "category_id": 1,
    "tag_ids": [1, 2],
    "bank": "Chase",
    "card": "Chase Freedom",
    "timestamp": "2024-01-15T10:30:00Z"
  }'
```

### Get Expenses with Filtering

```bash
curl -X GET "http://localhost:8080/api/v1/expenses?page=1&page_size=10&start_date=2024-01-01&end_date=2024-01-31" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Get Analytics by Category

```bash
curl -X GET "http://localhost:8080/api/v1/expenses/category?start_date=2024-01-01&end_date=2024-01-31" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 🔧 Configuration

### Environment Variables

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=finance
SSL_MODE=disable

# Authentication
JWT_SECRET=your-secret-key
GOOGLE_OAUTH_CLIENT_ID=your-google-client-id
```

## 🐳 Docker Deployment

### Using Docker Compose

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

## 🧪 Testing

### API Testing with Bruno

The project includes comprehensive API documentation and test collections using Bruno:

```bash
# Install Bruno (if not already installed)
npm install -g @usebruno/cli

# Run tests
bruno test
```

## 📊 Database Schema

### Core Tables

- **users** - User accounts and profiles
- **expenses** - Expense records with categories and tags
- **categories** - Expense categories
- **tags** - Expense tags for organization

### Key Features

- **Foreign Key Relationships** - Proper data integrity
- **Indexes** - Optimized query performance
- **Timestamps** - Automatic created_at/updated_at tracking
- **Soft Deletes** - Data preservation capabilities

## 🔒 Security Features

- **JWT Authentication** - Secure token-based authentication
- **CORS Support** - Cross-origin resource sharing
- **Input Validation** - Comprehensive request validation
- **SQL Injection Protection** - Parameterized queries
- **Environment-based Config** - Secure configuration management

## 🚀 Performance

- **Gin Framework** - High-performance HTTP web framework
- **GORM** - Efficient database operations
- **Connection Pooling** - Optimized database connections
- **Indexed Queries** - Fast data retrieval
- **Pagination** - Efficient large dataset handling

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

- 🐛 **Issues**: [GitHub Issues](https://github.com/lfelipessilva/finance-server/issues)

---

<div align="center">
  <p>Built with ❤️ using <a href="https://golang.org">Go</a>, <a href="https://gin-gonic.com">Gin</a>, and <a href="https://www.postgresql.org">PostgreSQL</a></p>
  <p>⭐ Star this repository if you found it helpful!</p>
</div>
