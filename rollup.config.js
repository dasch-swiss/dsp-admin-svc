import svelte from 'rollup-plugin-svelte';
import commonjs from '@rollup/plugin-commonjs';
import resolve from '@rollup/plugin-node-resolve';
import livereload from 'rollup-plugin-livereload';
import { terser } from 'rollup-plugin-terser';
import sveltePreprocess from 'svelte-preprocess';
import typescript from '@rollup/plugin-typescript';
import css from 'rollup-plugin-css-only';
import replace from "@rollup/plugin-replace";

const production = !process.env.ROLLUP_WATCH;

export default [
    {
        input: 'services/admin/frontend/main.ts',
        output: {
            sourcemap: true,
            format: 'iife',
            name: 'app',
            file: 'public/admin/build/bundle.js'
        },
        plugins: [
            replace({
                'process.env.BASE_URL': production ? JSON.stringify('PLACEHOLDER') : JSON.stringify('http://localhost:3000/'),
            }),
            svelte({
                preprocess: sveltePreprocess({ sourceMap: !production }),
                compilerOptions: {
                    // enable run-time checks when not in production
                    dev: !production
                }
            }),
            // we'll extract any component CSS out into a separate file - better for performance
            css({ output: 'bundle.css' }),

            // If you have external dependencies installed from npm, you'll most likely need these plugins. In
            // some cases you'll need additional configuration - consult the documentation for details:
            // https://github.com/rollup/plugins/tree/master/packages/commonjs
            resolve({
                browser: true,
                dedupe: ['svelte', 'svelte/transition', 'svelte/internal']
            }),
            commonjs(),
            typescript({
                sourceMap: !production,
                inlineSources: !production
            }),

            // Watch the `public` directory and refresh the browser on changes when not in production
            !production && livereload('public'),

            // For production builds minify, remove comments and logs
            // production && terser({
            //     format: {
            //         comments: false
            //     },
            //     compress: {
            //         drop_console: true
            //     }
            // }),
        ],
        watch: {
            clearScreen: false
        }
    },
];
