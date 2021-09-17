pipeline{
    
    agent any
    
    stages{
        
        stage('Build'){
            
            
            steps{
                
                cleanWs()                                     // Clears the data before each build
                
                //sh 'pwd'                                      -- this was just for checking we are in the right path
                
                sh 'git clone "https://github.com/j-saurabh/go-bookstacks.git"'
                
                dir('go-bookstacks')
                {
                    sh 'pwd'
                
                
                    sh 'docker build -t my-go-app .'
            
                    sh 'docker run -d -p 5000:5000 my-go-app'
            
                }    
            }        
        }
        
        
    }
