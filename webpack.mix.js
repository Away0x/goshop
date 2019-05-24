let mix = require('laravel-mix');

mix
  .setPublicPath('public/');

mix
  .ts('resources/js/app.ts', 'public/js')
  .sass('resources/sass/app.scss', 'public/css')
  .version()
  .copyDirectory('resources/svg', 'public/svg')
