'use strict';
var gulp = require('gulp');
var sass = require('gulp-sass');
//var compass = require('gulp-compass');
var cjsx = require('gulp-cjsx');
var gutil = require('gulp-util');
var path = require('path');

var paths = {
	coffee: 'static/src/coffee/*.coffee',
	cjsx: 'static/src/cjsx/*.cjsx',
	scss: 'static/src/scss/*.scss'
};

// Coffeescript
gulp.task('coffee', function () {
	return gulp.src(paths.coffee)
		.pipe(cjsx({ bare: false }).on('error', gutil.log))
		.pipe(gulp.dest('static/js'));
});

// Coffeescript-React
gulp.task('cjsx', function () {
	return gulp.src(paths.cjsx)
		.pipe(cjsx({ bare: false }).on('error', gutil.log))
		.pipe(gulp.dest('static/jsx'));
});

// Scss
gulp.task('sass', function () {
	return gulp.src(paths.scss)
		.pipe(sass().on('error', gutil.log))
		.pipe(gulp.dest('static/css'));
});

// Compass
/*gulp.task('compass', function () {
	gulp.src('./static/src/scss/*.scss')
		.pipe(compass({
			css: 'static/css',
			sass: 'static/src/scss'
		}))
		.on('error', gutil.log)
		.pipe(gulp.dest('static/css'));
});*/

gulp.task('build', ['coffee', 'cjsx', 'sass']);
gulp.task('watch', function () {
	gulp.watch(paths.coffee, ['coffee']);
	gulp.watch(paths.cjsx, ['cjsx']);
	gulp.watch(paths.scss, ['sass']);
});
