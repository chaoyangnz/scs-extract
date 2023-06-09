var fs = require("fs"),
    path = require("path")

const paths = ['']

const folder = "D:/ets2/base/"

const whitelist = ['def', 'ui', 'material', 'vehicle', 'automat']

function walk(dir, callback) {
    const files = fs.readdirSync(dir)
    files.forEach((file) => {
        var filepath = path.join(dir, file)
        const stats = fs.statSync(filepath)
        if (stats.isDirectory()) {
            const p = filepath.replaceAll('\\', '/').replace(folder, '')
            if(whitelist.includes(p.split('/')[0])) {
                paths.push(p)
            }

            walk(filepath, callback)
        } else if (stats.isFile()) {
            callback(filepath, stats)
        }
    })
}


walk(folder, () => {})

console.log()

fs.writeFileSync('./scs/well-known-paths.go', `package scs

var wellKnownPaths = []string{
${paths.map(p => '    "' + p + '"').join(', \n')},
}`)