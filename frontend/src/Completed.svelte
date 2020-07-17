<script>
    export let cardColumns;
    export let cardPromise;
    export let openTask;

    function handleModal() {
        window.$('#completed.ui.modal').modal('show')
    }
</script>


<h1 class="ui header">
  Completed Cards
</h1>
<div class="ui divider"></div>

{#await cardPromise}
  <div class="ui segment">
    <div class="ui active inverted dimmer">
      <div class="ui text loader">Loading</div>
    </div>
  </div>
{:then msg}
  <div class="ui {cardColumns(msg.cards.length)} cards">
    {#each msg.cards as card}
        {#if card.status === 1}
            <a class="ui card" href="/" on:click|preventDefault={openTask(card)}>
                <div class="content">
                <h1 class="header">{card.title}</h1>
                <div class="description">
                    {card.content}
                </div>
                </div>
            </a>
        {/if}
    {/each}
  </div>
{:catch error}
  <h1>{error}</h1>
{/await}