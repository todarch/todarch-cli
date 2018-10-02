pipeline {
  agent any

  stages {
    stage('Build') {
      steps {
        sh 'go build -o $GOPATH/bin/todarch'
      }
    }
    stage('Test') {
      steps {
        sh 'go test'
      }
    }
    stage('Release') {
      when {
        branch '0.0.x'
      }
      steps {
        sh "${env.JENKINS_SCRIPTS}/release-todarch.sh"
      }
    }
  }
}
