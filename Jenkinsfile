pipeline {
  agent any

  stages {

    stage('Build'){
      agent {
        docker {
          image 'golang:1.22.5-alpine'
          reuseNode true
        }
      }

      steps {
        sh '''
          echo '------------On Build Stage---------------'
          go version
          go run .
          echo '------------First POC Fin----------------'
        '''
      }
    }
  }
}
