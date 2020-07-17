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
            user: "eccc6483-e77b-4584-8ca8-3a434635d917"
        });
    }

    let finalize = async function(userID, cardID) {
        createService.finalize({userID: "eccc6483-e77b-4584-8ca8-3a434635d917", cardID: cardID});
    };

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
                        {#if card.status === 1}
                            <button button class="ui blue button" style="float: left;" type="submit" on:click={finalize("", card.id)}>Finalize</button>
                        {:else}
                            <button button class="ui blue disabled button" style="float: left;" type="submit">Finalize</button>
                        {/if}
                    {/if}
                {:catch err}
                    <h1>{err}</h1>
                {/await}
            {:else}
                {#await usersCardID}
                    <div class="ui segment">
                        <div class="ui active inverted dimmer">
                            <!-- <div class="ui text loader">Loading</div> -->
                        </div>
                    </div>
                {:then usersCardID}
                    {#if usersCardID.card.content === ""}
                        <button class="ui positive button" type="submit" on:click={volunteer("", card.id)}>Volunteer</button>
                    {:else}
                        <button class="ui positive disabled button" type="submit">Volunteer</button>
                    {/if}
                {/await}
            {/if}
            <button class="ui negative button">Cancel</button>
        </div>
    </div>
{/if}
