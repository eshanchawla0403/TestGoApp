// Run on an agent where we want to use Go
node {
    // Ensure the desired Go version is installed
    def root = tool type: 'go', name: '1.17.6'
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        stage('Preparation') { // for display purposes
            // Get some code from a GitHub repository
            git 'https://github.com/eshanchawla0403/TestGoApp.git'
        }
    
        stage('Build') {
            // Export environment variables pointing to the directory where Go was installed
            sh 'go version'
            sh 'go run main.go'
            sh 'go test -v ./...'
        }
    }    
}
