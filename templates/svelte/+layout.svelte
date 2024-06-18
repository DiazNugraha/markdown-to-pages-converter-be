<script>
  import "./styles.css";
</script>

<div class="container">
  <div class="navigation-wrapper">
    <div class="icon-wrapper">
      <div class="icon-frame"></div>
    </div>
    <div class="content-navigation">
      <!-- navigations -->    
    </div>
  </div>
  <main class="content-wrapper">
    <slot />
  </main>
</div>

<style>
  .container {
    width: 100%;
    height: 100vh;
    display: flex;
  }

  .icon-wrapper {
    position: -webkit-sticky; /* for Safari */
    position: sticky;
    top: 0;
    display: flex;
    padding: 10px;
    justify-content: center;
    align-items: center;
    backdrop-filter: blur(2px);
  }

  .icon-frame {
    width: 200px;
    height: 150px;
    background-color: white;
    border-radius: 20%;
  }

  .navigation-wrapper {
    flex: 2;
    background-color: #b0b0b0;
    overflow-y: scroll;
    padding: 10px 20px;
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .content-navigation {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .title {
    font-weight: bold;
    margin: 0;
    padding: 0;
    font-size: 20px;
  }

  .content-name {
    margin: 0;
    font-size: 15px;
    margin-left: 30px;
    text-decoration: none;
    color: black;
    padding: 5px 10px;
  }

  .content-name:hover {
    text-decoration: underline;
    background-color: skyblue;
  }

  .content-wrapper {
    flex: 8;
    background-color: #a0a0a0;
    padding: 10px 20px;
    overflow-y: scroll;
  }
</style>
