<?php

/*
|--------------------------------------------------------------------------
| Application Routes
|--------------------------------------------------------------------------
|
| Here is where you can register all of the routes for an application.
| It's a breeze. Simply tell Laravel the URIs it should respond to
| and give it the controller to call when that URI is requested.
|
*/

Route::get('/', function () {
    return view('welcome');
});


Route::get('wel/{id}/{name}', function ($id, $name) {
    return 'wel ' . $id . ' ' . $name;
});


Route::get('posts/{post}/comments/{comment}', function ($post, $comment) {
    return $post . ' ' . $comment;
});


Route::get('user_{id}', function ($id) {
    return $id;
});
