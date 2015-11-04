module.exports = function(grunt) {
    'use strict';

    // var config = grunt.file.readJSON('package.json'),
    var config = {
        src: './src',
        dist: './dist',
        tmp: './.tmp'
    };

    grunt.initConfig({
        
        config: config,

        typescript: {
            base: {
                src: ['<%= config.src %>/**/*.ts'],
                dest: '<%= config.dist %>',
                options: {
                    // module: 'amd', //or commonjs 
                    // target: 'es5', //or es3 
                    // basePath: 'path/to/typescript/files',
                    sourceMap: true,
                    declaration: true
                }
            }
        }
    });

    grunt.loadNpmTasks('grunt-typescript');

    grunt.registerTask('run', 'Compile and run the app', function() {
        grunt.task.run([
            'typescript'
        ]);
    });
};
