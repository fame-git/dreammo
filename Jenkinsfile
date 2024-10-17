pipeline {
  agent any

  stages {

    stage('Verify Agent'){
      agent {
        docker {
          image 'go-agent'
          reuseNode true
        }
      }

      steps {
        sh '''
          go version

          # Set GOCACHE to a directory with appropriate permissions
          export GOCACHE=/tmp/.cache
          mkdir -p /tmp/.cache

          echo "------fin check go--------------"

          kubectl version --client
          echo "------fin kubectl verify--------"

          
        '''
      }
    }
  }
}
