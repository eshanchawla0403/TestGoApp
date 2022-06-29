node {
    // Go version
    def root = tool type: 'go', name: '1.17.6'
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        stage('Build') {
            // Get code
            git 'https://github.com/eshanchawla0403/TestGoApp.git'
        }
    
        stage('Run&Test') {
            sh 'go run main.go'
            sh 'go test -v ./...'
        }
    }    
}
