FRONTEND_DIR := "frontend"
BACKEND_DIR := "backend"
WASM_DIR := "frontend/dist"
WASM_OUTPUT := "frontend/dist/main.wasm"
STATIC_DIR := "frontend/static"
BACKEND_STATIC_DIR := "backend/cmd/server/static"
BACKEND_BINARY := "backend/bin/server"

CHROME_DEV_PORT := "9222"

#Build tailwind styles
# requires tailwind cli (pnpm i -g tailwindcss)
tailwind:
    tailwindcss -i frontend/static/input.css -o frontend/static/style.css
alias tw := tailwind

#Run chrome with developer port
# requires chrome, duh
chrome:
    @google-chrome --remote-debugging-port={{CHROME_DEV_PORT}} --user-data-dir=/tmp/chrome-dev-profile --window-size=1920x1080 --new-window "http://localhost:8080" &> /dev/null &

#Reload chrome webpage
# requires websocat (cargo binstall websocat) and jq (https://jqlang.github.io/jq/)
reload-chrome:
    #!/usr/bin/env bash
    {
        sleep 2
        echo '{"id":1,"method":"Page.reload","params":{}}' | websocat $(curl -s http://localhost:9222/json | jq -r '.[0].webSocketDebuggerUrl')
    } &> /dev/null || true
    echo -e "\e[32mTriggered reload\e[0m"
    exit 0

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
# requires reflex (go install github.com/cespare/reflex@latest)
watch:
    reflex -r '(^{{FRONTEND_DIR}}/static)|(\.go$$)' -s -- sh -c 'just reload-chrome & just c run-backend'
alias w := watch

# Clean
clean:
    rm -rf {{BACKEND_BINARY}} {{BACKEND_STATIC_DIR}} {{WASM_OUTPUT}}
alias c := clean
