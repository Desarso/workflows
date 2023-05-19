When thinking about the most useful coding things I have done, byt far, my assembly make file project setUp and automation was the most useful, for this reason I have started this project to think about and code some workflow automations, sincet this will set me apart from others, and will help me in the long run.


First things first. 
I think that making a command alias for setting up a solidjs project would be good, I think that I should create my own base project, and include all the things I would want to include, and then make a command that clones and runs yarn install.

Some things I might want to include by default:
    Tailwinds, CSS, my SASS structure for files,
    Secondly, I should develope a desing architechture with SASS where each component has an ID,
    and I only style withing that Id when designing a component so that I don't get styling clonflicts in the long run. I can also transfer stlying that is particular to one element. 


Let's think about full deployment, of an app. 
I should have a testing url website, where when I deploy to the test branch it then puts that page withing the 
nginx server.


So let's be clear, I am not going to automate any of the changes, to the nginx server, only the deplyments of the projects connected to their respective urls. 


So when I am working on the chessbot, I want to be working int the production branch, then I to be able to merge with the test branch, then I should create a workflow in github that grabs that branch, and then it needs to run a couple of things, it should build the project, using yarn build, then it should copy the build folder to the nginx server, to the correct part. 


//Ok so instead of using anything else, I am gonna make it such that the base solidjs folder, pushed the build. 
//Then I am gonna have a make file deploy command, and a make file deploy to production. They will run different commands, the production one should prompt for confirmation.

//the way it is gonna happend, I am gonna have a go microservice in the backend, it will receive a deply command, that will yarn build, then gitpush, then it will send an http request to the microservice, this gets routed thru nginx. Into a go program, which will git pull, certain files, I might need a better way to do this. Anyways it should fetch only the build file, and then I it will put it in the right place. problem is, that I need to be able to do this generally, in all scenerarios, so I need to be able to do this for all projects.

So the microservice should receice the actual github repo. 


In other news, I also want an automated way to setup an nginx repo fast. 