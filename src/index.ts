class FileClass {
    fullpath: string;
    
    constructor(fullpath: string) {
        this.fullpath = fullpath;
    }
}

var file1 = new FileClass('testpath');
console.log(file1.fullpath);
