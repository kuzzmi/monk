/// <reference path="../typings/node/node.d.ts" />

var fs = require('fs');

class FileClass {
    constructor(public fullpath: string) {
    }
}

class Folder {
    files:  Array<FileClass> = [];

    constructor(public path:string, public name:string, files: Array<string>) {
        for(var index in files) {
            this.files.push(new FileClass(files[index]));
        }
    }
}

fs.readdir('.', function(err, data) {
    var folder = new Folder(__dirname, 'monk', data);

    console.log(folder);
});

var file1 = new FileClass('testpath');
console.log(file1.fullpath);
