pipeline {
    agent {
        docker { image 'izabelacg/ruby-wf-cli:latest' }
    }
    parameters {
        string(name: 'WF_ENDPOINT', defaultValue: 'wavefront.vmware.com', description: 'Wavefront Instance URL')
        string(name: 'WF_PROXY', defaultValue: '', description: 'Wavefront Proxy')
        string(name: 'METRIC_NAME', defaultValue: 'tanzu.vmware.com/serverless', description: '')
    }
    stages {
        stage('search-wave') {
            environment {
                WAVEFRONT_ENDPOINT = "${param.$WF_ENDPOINT}
                WAVEFRONT_PROXY = "${param.WF_PROXY}"
                WAVEFRONT_TOKEN = credentials('wf_token')
                METRIC = "${param.METRIC_NAME}"
                METRIC_PREFIX = ''
            }

            steps {
                sh './scripts/wf-script'
            }
        }
    }
}