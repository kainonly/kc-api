<?php

namespace App\Http\Controllers\Erp;

use Illuminate\Http\Request;
use lumen\curd\CurdController;

class Base extends CurdController
{
    public function __construct(Request $request)
    {
        parent::__construct($request);
    }
}