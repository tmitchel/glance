<script>
    export let cardColumns;
    export let cardPromise;
    export let openTask;

    function handleModal() {
        window.$('.ui.modal').modal('show')
    }

</script>


<h1 class="ui header">
  Unclaimed Cards
</h1>
<div class="ui divider"></div>

<button class="ui compact labeled icon button"  on:click={handleModal}>
  <i class="plus icon fitted"></i>New Task
</button>

{#await cardPromise}
  <div class="ui segment">
    <div class="ui active inverted dimmer">
      <div class="ui text loader">Loading</div>
    </div>
  </div>
{:then msg}
  <div class="ui {cardColumns(msg.cards.length)} cards">
    {#each msg.cards as card}
      <a class="ui card" href="/" on:click|preventDefault={openTask(card)}>
        <div class="content">
          <h1 class="header">{card.title}</h1>
          <div class="description">
            {card.content}
          </div>
        </div>
      </a>
    {/each}
  </div>
{:catch error}
  <h1>{error}</h1>
{/await}