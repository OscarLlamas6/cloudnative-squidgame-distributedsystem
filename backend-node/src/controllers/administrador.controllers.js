//const db = require('../database')
const estudianteControllers = {};


//###########################################
//####### END POINT 1 #######################
//###########################################

estudianteControllers.test = async (req, res) => {
    return res.status(200).send({ results: 'exito' })
}

module.exports = estudianteControllers