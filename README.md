# go-bookstacks

A simple golang service that I have written for learning more about golang. You can add books either you have read, or TBR. 
It is sort of a basic as of now. But I will keep adding to it. 

It has 2 fields : book name and author and it lists the already added books on the top. 


To see how this goes, run ```$go build``` in the folder where all the code files of this repository are cloned. 

The $go build will create a binary with the name of the module. (Say, bookstacks)

Then, run ./bookstacks and check for the output in ```localhost:5000/assets/```

<h2> Added Dockerfile to the project </h2>

Once the project is cloned, you can run the dockerfile directly. 

Run the below commands to run the project through dockerfile.

```docker build -t my-go-app .``` </br>
```docker run -d -p 5000:5000 my-go-app```

Then, go to ```localhost:5000/assets/``` to check the output. 

<h2> Added Jenkinsfile </h2>

Jenkinsfile which has the docker commands will run and check if the build succeeds.
First, it clones this repository and then does the dockerizing part. 


Create a pull request or an issue if there is anything you can contribute to this project.
