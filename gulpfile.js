'use strict';
var gulp = require('gulp');
var sass = require('gulp-sass');
var riot = require('gulp-riot');
var gutil = require('gulp-util');
var coffee = require('gulp-coffee');
var sourcemaps = require('gulp-sourcemaps');

var paths = {
	coffee: 'static/src/coffee/*.coffee',
	riot: 'static/src/riot/*.tag',
	scss: 'static/src/scss/*.scss'
};

// Coffeescript
gulp.task('coffee', function () {
	return gulp.src(paths.coffee)
		.pipe(sourcemaps.init())
		.pipe(coffee({ bare: false }).on('error', gutil.log))
		.pipe(sourcemaps.write())
		.pipe(gulp.dest('static/js'));
});

gulp.task('riot', function () {
	return gulp.src(paths.riot)
		.pipe(riot({type: "coffee"}))
		.pipe(gulp.dest('static/riot'));
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

gulp.task('build', ['coffee', 'riot', 'sass']);
// FIXME: crashes on error!
gulp.task('watch', function () {
	gulp.watch(paths.coffee, ['coffee']).on('error', gutil.log);
	gulp.watch(paths.riot, ['riot']).on('error', gutil.log);
	gulp.watch(paths.scss, ['sass']).on('error', gutil.log);
});
