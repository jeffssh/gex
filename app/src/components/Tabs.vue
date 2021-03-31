<template>
    <div class="Tabs">
      <div class="tabs">
        <ul>
          <li v-for="tab in tabs" :class="{ 'is-active': tab.isActive }" v-bind:key=tab.name>
            <a :href="tab.href" @click="selectTab(tab)">{{ tab.name }}</a>
          </li>
        </ul>
      </div>

      <div class="tabs-details"> 
        <slot></slot>
      </div>
    </div>
</template>

<script>
export default {
  name: 'Tabs',
  data() {
    return { tabs: [], useChildren: {default: true} }
  },
  methods: {
    selectTab: function(selectedTab) {
      this.tabs.forEach(tab => {
        tab.isActive = (tab.name == selectedTab.name);
      });
    }
  },
  mounted() {
    //console.log('in tabs mount')
    if (this.useChildren) {
      this.tabs = this.$children;
    }
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.section {
  padding: 2em 0;
}
</style>
