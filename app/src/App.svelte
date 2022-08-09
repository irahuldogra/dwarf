<script lang="ts">
  import List from './components/List.svelte';
  import New from './components/New.svelte';
</script>

<main>
  <header>
    <h1 class="header">Dwarf -- Go Url Shortener</h1>
    <New />
  </header>
  <List />
</main>

<style>
  :root {
    font-size: 16px;
    line-height: 24px;
    font-weight: 400;
    font-family: 'Open Sans', sans-serif;
    color-scheme: light dark;
    color: rgba(255, 255, 255, 0.87);
    background-color: #242424;

    font-synthesis: none;
    text-rendering: optimizeLegibility;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    -webkit-text-size-adjust: 100%;
  }

  header {
    display: flex;
    align-items: center;
  }

  .header {
    margin-left: 20px;
    margin-right: 20px;
    letter-spacing: 2px;
    text-decoration: none;
  }
</style>
