FRONTEND_DIR := "frontend"
BACKEND_DIR := "backend"
WASM_DIR := "frontend/dist"
WASM_OUTPUT := "frontend/dist/main.wasm"
STATIC_DIR := "frontend/static"
BACKEND_STATIC_DIR := "backend/cmd/server/static"
BACKEND_BINARY := "backend/bin/server"

#Build tailwind styles
tailwind:
    tailwindcss -i frontend/static/input.css -o frontend/static/style.css
alias tw := tailwind

# Build frontend
frontend: tailwind
    cd {{FRONTEND_DIR}} && GOOS=js GOARCH=wasm go build -o ../{{WASM_OUTPUT}} ./cmd/wasm
alias f := frontend

# Run frontend
run-frontend: frontend
    cp -r {{STATIC_DIR}}/* {{WASM_DIR}}
    cd {{FRONTEND_DIR}} && miniserve --index index.html dist
alias rf := run-frontend

# Build backend
backend: frontend
    cd {{BACKEND_DIR}}/cmd/server && go generate && go build -o ../../bin/server
alias b := backend

# Run backend
run-backend: frontend
    cd {{BACKEND_DIR}}/cmd/server && go generate && go run .
alias r := run-backend

# Run backend continuously
watch:
    reflex -r '(^{{FRONTEND_DIR}}/static)|(\.go$$)' -s just run-backend
alias w := watch

# Clean
clean:
    rm -rf {{BACKEND_BINARY}} {{BACKEND_STATIC_DIR}} {{WASM_OUTPUT}}
alias c := clean
