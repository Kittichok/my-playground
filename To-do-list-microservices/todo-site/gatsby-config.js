require("dotenv").config({
  path: `.env.${process.env.NODE_ENV}`,
})

module.exports = {
  // siteMetadata: {
  //   title: "TodoSite",
  // },
  plugins: [
    "gatsby-plugin-styled-components", 
    {
      resolve: `gatsby-plugin-typography`,
      options: {
        pathToConfigModule: `src/utils/typography`,
      },
    },
    {
      resolve: `gatsby-plugin-create-client-paths`,
      options: { prefixes: [`/app/*`] },
    },
    'gatsby-plugin-postcss',
  ],
};
