var gulp    = require('gulp'),
    gogo    = require("gulp-go");


var server;

gulp.task('run', function() {
  server = gogo.run('main.go', [], {cwd: __dirname, stdio: 'inherit'});
});
gulp.task('watch', ['run'], function() {
  gulp.watch('**/*.go').on("change", function() {
    if (server) server.restart();
  });
});

gulp.task('default', ['watch']);
