

<template>
  <div id="app">
    <button id="dark-button" class="button is-dark" style="margin-left: 10px;" :userObj="data" v-on:click="updateData('AAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDD')">Send Packet</button>
    <Tabs>
      <Tab name="TCP Intercept" :selected="true">
        <div style="height: 50px;">
          <!-- https://bulma.io/documentation/elements/button/ -->
          <buttons>
            <button id="dark-button" class="button is-dark" style="margin-left: 10px;">Forward</button>
            <button id="dark-button" class="button is-dark" style="margin-left: 10px;">Drop</button>
            <button id="dark-button" class="button is-dark" style="margin-left: 10px;">Intercept is off</button>
          </buttons>
        </div>
          <Tabs>
            <Tab name="Hex Editor" :selected="true">
              <div>
                <HexEditor :data="initdata"/>
              </div>
            </Tab>
            <Tab name="Text Editor">
              <TextEditor v-on:packet-modified="updateData($event)" :data="initdata"/> 
            </Tab>
            <Tab name="Settings">
            </Tab>
          </Tabs>
      </Tab>
      <Tab name="TCP Proxy"></Tab>
      <Tab name="TCP Repeater"></Tab>
    </Tabs>
  </div>
</template>

<script>
import Vue from 'vue'
import TextEditor from './components/TextEditor.vue'
import HexEditor from './components/HexEditor.vue'
//import Table from './components/Table.vue'

import Tab from './components/Tab.vue'
import Tabs from './components/Tabs.vue'

export default {
  name: 'App',
  components: {
    TextEditor,
    HexEditor,
    Tab,
    Tabs
  },
    data() {
    return { initdata: "\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC" }
  },
  methods: {
    say: function (msg) {
      alert(msg)
    },
    updateData: function (data) {
      console.log(data)
      this.initdata = data
      //$refs.InterceptTextEditor.loadData(data)
    }
  }
}

</script>

<style>
@import "css/bulma.min.css";

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  overflow: auto;
}

#dark-button {
  float: left;
  padding: 10px
}

</style>
