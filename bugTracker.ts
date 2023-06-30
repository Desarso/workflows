

let apps: any = [];

class Bug{
    private appName: string;
    private production: boolean;
    private error: string;
    private createdAt: Date;
    private status = "open";
    constructor(appName: string, production: boolean, error: string){
        this.appName = appName;
        this.production = production;
        this.error = error;
        this.createdAt = new Date();
        //I must check if app exists, if it does not, I must create it.
        //and add bug to it
        for(let i = 0; i < apps.length; i++){
            if(apps[i].getAppName() == appName){
                apps[i].addBug(this);
                return;
            }
        }
        let app = new App(appName);
        app.addBug(this);
    }
    getError(){
        return this.error;
    }
    getCreatedAt(){
        return this.createdAt;
    }
    getProduction(){
        return this.production;
    }
    getAppName(){
        return this.appName;
    }

    resolve(){
        //call parent app, and tell it to resolve this bug.
        this.status = "closed";

    }
}


class App{
    //app are instantiated, when a bug is created, and it belongs to an app that does not exist yet.
    private bugs: Bug[];
    private appName: string;
    private appID: string;
    constructor(appName: string){
        this.appName = appName;
        this.appID = this.generateID();
        this.bugs = [];
        apps.push(this);
    }


    generateID(){
        //make id 20 characters long
        let id = "";
        let possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
        for(let i = 0; i < 20; i++){
            id += possible.charAt(Math.floor(Math.random() * possible.length));
        }
        return id;
    }

    addBug(bug: Bug){
        this.bugs.push(bug);
    }

    getAppName(){
        return this.appName;
    }


}


function updateApp(){
    //best way to do this is to store everything as json.
    //first I get the file
    //then I parse it
    //then update the apps array.
}





//let's think about but workflow, bug happends in the app tester it needs to be created, bug will grab data from the current env.
//based on the info it will create the bug. for this reason the bug should be independent from the app.
//however on creation, the bug should tell you what app it is from, and it's point of origin. Since this is the case, 
//app will be used, so I should try to keep some consistent identifier betwee, production, and development, apps. 
//we have the classed, but I am currently working with an array of classes in the actual program, instead of doing this, and referencing,
//the array as a proxy, I want to get the array from the database.
