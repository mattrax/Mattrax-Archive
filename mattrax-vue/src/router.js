import Vue from "vue";
import Router from "vue-router";

//Views
import Dashboard from "./views/Dashboard.vue";
import NotFound from "./views/NotFound.vue";

Vue.use(Router);
export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "home",
      component: Dashboard
    },
    {
      path: '*',
      component: NotFound //TODO: Lazy Loading For Preformace
    }
    /*{
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () =>
        import( "./views/About.vue") //webpackChunkName: "about" //Use This ChunkName TO Name The Output File
    }*/
  ]
});
