<template>
  <div class="Repeater">
   <div style="height: 50px;" ref="repeater">
      <buttons>
        <button id="dark-button" class="button is-dark" style="margin-left: 10px;">Send</button>
      </buttons>
    </div>
    
  <ejs-tab id='dynamic_tab' ref='TabInstance'  :selected= "tabSelected">
    <e-tabitems>    
    </e-tabitems>
  </ejs-tab>

    <h1>test</h1>
  </div>
</template>

<script>
import Tabs from './Tabs.vue'
import Tab from './Tab.vue'
import Vue from 'vue'

import { TabPlugin } from '@syncfusion/ej2-vue-navigations';
import { createElement } from '@syncfusion/ej2-base';

Vue.use(TabPlugin);

export default {
  name: 'Repeater',
  components: {
  },
  props: {
    repeaterTabs: Object
  },
  mounted() {
  // https://ej2.syncfusion.com/vue/documentation/tab/how-to/load-tab-items-dynamically/
    var tabObj = this.$refs.TabInstance.ej2Instances;
    //console.log('repeater tabs', this.repeaterTabs)
    let totalItems = document.querySelectorAll('#dynamic_tab .e-toolbar-item').length;
    console.log(this.repeaterTabs.keys())
    this.repeaterTabs.forEach(function(v, k) {
        console.log(k)
        let item =  { header: { text: k.toString() }, content: v };
        tabObj.addTab([item], (totalItems++));
      }
    )
  },
  data: function(){
    return {

      headerText0: { text: 'ASP.NET' },
      headerText1: { text: 'ASP.NET MVC' },
      headerText2: { text: 'JavaScript' },

    content0: 'ASP.NET is an open-source server-side web application framework designed for web development to produce ' +
    'dynamic web pages. It was developed by Microsoft to allow programmers to build dynamic web sites, web applications ' +
    'and web services. It was first released in January 2002 with version 1.0 of the .NET Framework, and is the successor ' +
    'to Microsoft\'s Active Server Pages (ASP) technology. ASP.NET is built on the Common Language Runtime (CLR), allowing ' +
    'programmers to write ASP.NET code using any supported .NET language. The ASP.NET SOAP extension framework allows ' +
    'ASP.NET components to process SOAP messages.',

     content1: 'The ASP.NET MVC is a web application framework developed by Microsoft, which implements the ' +
    'model–view–controller (MVC) pattern. It is open-source software, apart from the ASP.NET Web Forms component which is ' +
    'proprietary. In the later versions of ASP.NET, ASP.NET MVC, ASP.NET Web API, and ASP.NET Web Pages (a platform using ' +
    'only Razor pages) will merge into a unified MVC 6.The project is called ASP.NET vNext.',

     content2: 'JavaScript (JS) is an interpreted computer programming language. It was originally implemented as ' +
    'part of web browsers so that client-side scripts could interact with the user, control the browser, communicate ' +
    'asynchronously, and alter the document content that was displayed.[5] More recently, however, it has become common in ' +
    'both game development and the creation of desktop applications.',

    }
   }
}


</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  @import "../../node_modules/@syncfusion/ej2-base/styles/material.css";
  @import "../../node_modules/@syncfusion/ej2-buttons/styles/material.css";
  @import "../../node_modules/@syncfusion/ej2-popups/styles/material.css";
  @import "../../node_modules/@syncfusion/ej2-vue-navigations/styles/material.css";
</style>
