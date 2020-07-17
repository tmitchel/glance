<script>
    import Drop from "./Drop.svelte";

    export let card;
    export let usersCardID;
    export let volunteer;
    export let createService;

    let status = 0;

    let updateStatus = async function(cardID) {
        createService.updateStatus({
            status: parseInt(window.$('.ui.dropdown').dropdown('get value')),
            card: cardID,
            user: "888aabf9-6289-4b74-afc6-976a4a74ad4a"
        });
    }

    const statuses = [
        "Not Started",
        "Completed",
        "Help Requested",
        "In Progress",
    ];
</script>

{#if card}
    <div id="task" class="ui modal">
        <div class="content">
            <h1 class="header">{card.title}</h1>
            <p>{card.content}</p>
        </div>
        <div class="actions">
            {#if card.claimed}
                {#await usersCardID}
                    <div class="ui segment">
                        <div class="ui active inverted dimmer">
                            <!-- <div class="ui text loader">Loading</div> -->
                        </div>
                    </div>
                {:then usersCardID}
                    {#if card.id === usersCardID.card.id}
                        <svelte:component this={Drop} statuses={statuses} card={card}></svelte:component>
                        <button button class="ui positive button" type="submit" on:click={updateStatus(card.id)}>Submit</button>
                    {/if}
                {:catch err}
                    <h1>{err}</h1>
                {/await}
            {:else}
                <button class="ui positive button" type="submit" on:click={volunteer("", card.id)}>Volunteer</button>
            {/if}
            <button class="ui negative button">Cancel</button>
        </div>
    </div>
{/if}
