const childProcess = require('child_process');
const os = require('os');
const gulp = require('gulp');
const rename = require('gulp-rename');
const chalk = require('chalk');

function buildProject() {
  return new Promise((resolve, reject) => {
    const child = childProcess.spawn('npm', ['run', os.platform() === 'win32' ? 'build:win' : 'build:unix']);

    child.stdout.on('data', (data) => {
      console.log(data.toString());
    });

    child.stderr.on('data', (data) => {
      console.log(chalk.red(data.toString()));
    });

    child.on('close', () => {
      resolve();
    });

    child.on('error', (err) => {
      console.error(err);
      reject(err);
    });
  });
}

/* 拷贝template */
function copyTemplate() {
  return gulp.src('template/**/*.*')
    .pipe(gulp.dest('dist/github-cli/template'))
}

/* 拷贝配置文件 */
function copyConfigFile() {
  return gulp.src('config-example.json')
    .pipe(rename(function(p) {
      p.basename = 'config';
    }))
    .pipe(gulp.dest('dist/github-cli'))
}

/* 拷贝许可证 */
function copyLicense() {
  return gulp.src('LICENSE')
    .pipe(gulp.dest('dist/github-cli'))
}

exports.default = gulp.parallel(buildProject, copyTemplate, copyConfigFile, copyLicense);