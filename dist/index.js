var FileClass = (function () {
    function FileClass(fullpath) {
        this.fullpath = fullpath;
    }
    return FileClass;
})();
var file1 = new FileClass('testpath');
console.log(file1.fullpath);
//# sourceMappingURL=index.js.map