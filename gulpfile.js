'use strict';
var gulp = require('gulp');
var sass = require('gulp-sass');
//var compass = require('gulp-compass');
var cjsx = require('gulp-cjsx');
var gutil = require('gulp-util');
var path = require('path');

var paths = {
	coffee: 'static/src/coffee/*.{cjsx,coffee}',
	scss: 'static/src/scss/*.scss'
};

// Coffeescript-React
gulp.task('cjsx', function () {
	return gulp.src(paths.coffee)
		.pipe(cjsx({ bare: false }).on('error', gutil.log))
		.pipe(gulp.dest('static/js'));
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

gulp.task('watch', function () {
	gulp.watch(paths.coffee, ['cjsx']);
	gulp.watch(paths.scss, ['sass']);
});
