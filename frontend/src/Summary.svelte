<script>
    export let cardColumns;
    export let userPromise;
    export let color;
    export let openTask;
</script>

<div>
  {#await userPromise}
    <div class="ui segment">
      <div class="ui active inverted dimmer">
        <div class="ui text loader">Loading</div>
      </div>
    </div>
  {:then msg}
    <div class="ui {cardColumns(msg.pairs.length)} cards">
      {#each msg.pairs as pair}
        <div class="ui card">
          <div class="content">
            <h1 class="header">{pair.User.name}</h1>
            {#if pair.Card.content !== ""}
              <a class="ui card" href="/" on:click|preventDefault={openTask(pair.Card)}>
                <div class="{color(pair.Card.status)} content">
                  <h3 class="custom header">{pair.Card.title}</h3>
                  <div class="description">
                    {pair.Card.content}
                  </div>
                </div>
              </a>
            {:else}
              <h1 class="ui header centered">No task</h1>
            {/if}
          </div>
        </div>
      {/each}
    </div>
  {:catch error}
    <h1>{error}</h1>
  {/await}
</div>

<style>
.blue-bkg{
  background-color: #2185D0 !important;
}
.green-bkg{
  background-color: #21BA45 !important;
}
.red-bkg{
  background-color: #DB2828 !important;
}
</style>