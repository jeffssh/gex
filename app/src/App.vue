

<template>
  <div id="app">
    <!--<TextEditor/>-->
    <HexEditor/>
  <!--
    <Tabs :tabs='
    [
      {
        title: "Pictures",
        content: "Pictures content"
      },
      {
        title: "Music",
        content: "Music content. Wanna see some <a href=\"#\" data-show=\"3\">Documents</a> content?"
      },
      {
        title: "Videos",
        content: "Videos content. <a href=\"#\" data-alert=\"VIDEOS!!!\">Alert videos</a>"
      },
      {
        title: "Documents",
        content: "Documents content. Wanna see some <a href=\"#\" data-show=\"1\">Music</a> content?"
      },   
    ]'/>
  -->

  </div>
</template>

<script>
import Vue from 'vue'
//import TextEditor from './components/TextEditor.vue'
import HexEditor from './components/HexEditor.vue'
//import Table from './components/Table.vue'

//import Tab from './components/Tab.vue'
//import Tabs from './components/Tabs.vue'

export default {
  name: 'App',
  components: {
    //TextEditor,
    HexEditor
    //Tab,
    //Tabs
  }
}

/*
var tabs = [
  {
    title: "Pictures",
    content: "Pictures content"
  },
  {
    title: "Music",
    content: "Music content. Wanna see some <a href=\"#\" data-show=\"3\">Documents</a> content?"
  },
  {
    title: "Videos",
    content: "Videos content. <a href=\"#\" data-alert=\"VIDEOS!!!\">Alert videos</a>"
  },
  {
    title: "Documents",
    content: "Documents content. Wanna see some <a href=\"#\" data-show=\"1\">Music</a> content?"
  },   
];
*/
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

</style>
