# Setup
#default_registry('docker.io/robjkc')
#allow_k8s_contexts(k8s_context())

# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')

local_resource(
    'go-hello-compile',
    cmd='CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/go-hello -gcflags "-N -l" ./',
    deps=['./main.go']
)

docker_build_with_restart(
    'go-hello-image',
    '.',
    dockerfile='./tilt/Dockerfile',
    #entrypoint=['/dlv', '--listen=:40000', '--api-version=2', '--headless=true', 'exec', '/app/build/go-hello'],
    entrypoint=['/app/build/go-hello'],
    only=[
        './build',
    ],
    live_update=[
        sync('./build', '/app/build'),
    ],
)

k8s_yaml('tilt/go-hello.yaml')

k8s_resource('go-hello', port_forwards=['8089:8082', '40000:40000'],
             resource_deps=['go-hello-compile'])
