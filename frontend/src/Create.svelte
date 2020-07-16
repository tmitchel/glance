<script>
    let cardTitle = "";
    let cardContent = "";
    export let createService;

    async function createTask() {
        createService.card({
            title: cardTitle,
            content: cardContent,
            Creator: "888aabf9-6289-4b74-afc6-976a4a74ad4a"
        });
    }
</script>

<div class="ui modal">
  <div class="header">New Task</div>
  <form class="ui form content">
    <div class="field">
      <label>Title</label>
      <input type="text" bind:value={cardTitle} placeholder="Task Title">
    </div>
    <div class="field">
      <label>Description</label>
      <textarea bind:value={cardContent} rows=4 placeholder="Describe the task..."></textarea>
    </div>
  </form>
  <div class="actions">
    <button class="ui positive button" type="submit" on:click={createTask}>Submit</button>
    <button class="ui negative button">Cancel</button>
  </div>
</div>