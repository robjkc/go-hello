default_registry('docker.io/robjkc')
allow_k8s_contexts(k8s_context())
docker_build('go-hello', '.', dockerfile='Dockerfile')
k8s_yaml('tilt/go-hello.yaml')
k8s_resource('go-hello', port_forwards='8089:8082')