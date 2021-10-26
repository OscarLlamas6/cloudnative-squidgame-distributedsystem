const { Router } = require('express');
const router = Router();

const administradorControllers = require('../controllers/administrador.controllers');

router.get('/test', administradorControllers.test);


module.exports = router;