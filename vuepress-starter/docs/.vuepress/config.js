module.exports = {
    title: 'YellowAndGreen',
    description: 'A few blogs and notes',
    dest: './dist',  // build的目录
    port: '7777',
    head: [
        ['link', {rel: 'icon', href: '/logo.jpg'}],
        ['link',{rel:'stylesheet',href:'/css/head_logo.css'}]
    ],
    markdown: {
        lineNumbers: true
    },
    themeConfig: {
        nav: require("./nav.js"),
        sidebar: require("./sidebar.js"),
        sidebarDepth: 2,
        lastUpdated: 'Last Updated',
        searchMaxSuggestoins: 10,
        serviceWorker: {
            updatePopup: {
                message: "有新的内容.",
                buttonText: '更新'
            }
        },
        editLinks: true,
        editLinkText: '在 GitHub 上编辑此页 ！'
    }
}
