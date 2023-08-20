/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./frontend/static/*.html"],
    theme: {
        extend: {
            fontFamily: {
                'serif': ['Soria', 'ui-serif', 'Georgia'],
                'mono': ['"0xProto"', 'ui-monospace', 'SFMono-Regular'],
            }
        },
    }
}
