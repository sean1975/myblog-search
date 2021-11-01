function initAutoComplete(inp) {
  let abortController;

  function getSearchSuggestions(keyword) {
    // js is single-thread executed, so no needs for accessing abortController in sync mode
    if (abortController) {
      abortController.abort();
    }
    if (!keyword) {
      removeAutoCompleteItems();
      return;
    }
    // AbortController instances cannot be reused and need to re-created
    abortController = new AbortController();
    fetch(
      '/autocomplete/?query=' + keyword, {
        signal: abortController.signal
      }
    ).then(
      response => response.json()
    ).then(
      results => {
        if (results.length > 0) {
          createAutoCompleteItems(results);
        } else {
          removeAutoCompleteItems();
        }
      }
    ).catch(
      err => {
        if (err.name != 'AbortError') {
          throw err;
        }
      }
    ).finally(
      () => {
        abortController = null;
      }
    );
  }

  function getAutoCompleteList() {
    return document.getElementById("autocomplete-list");
  }

  function getAutoCompleteItems() {
    let autoCompleteList = getAutoCompleteList();
    if (autoCompleteList) {
      return autoCompleteList.getElementsByTagName("DIV");
    } else {
      return null;
    }
  }

  function getActiveAutoCompleteItem() {
    let autoCompleteList = getAutoCompleteList();
    if (autoCompleteList) {
      let activeAutoCompleteItems = autoCompleteList.getElementsByClassName("autocomplete-active");
      if (activeAutoCompleteItems) {
        return activeAutoCompleteItems.item(0);
      }
    }
  }

  function createAutoCompleteItems(searchSuggestions) {
    let autoCompleteList = getAutoCompleteList();
    if (autoCompleteList) {
      autoCompleteList.textContent = '';
      for (i = 0; i < searchSuggestions.length; i++) {
        let autoCompleteItem = document.createElement("DIV");
        autoCompleteItemLink = document.createElement('A');
        autoCompleteItemText = document.createTextNode(searchSuggestions[i].Title);
        autoCompleteItemLink.appendChild(autoCompleteItemText);
        autoCompleteItemLink.title = searchSuggestions[i].Title;
        autoCompleteItemLink.href = searchSuggestions[i].Url;
        autoCompleteItem.appendChild(autoCompleteItemLink);
        autoCompleteList.appendChild(autoCompleteItem);
      }
    }
  }

  function removeAutoCompleteItems() {
    let autoCompleteList = getAutoCompleteList();
    if (autoCompleteList) {
      autoCompleteList.textContent = '';
    }
  }

  let currentFocus;
  // when input value has changed
  inp.addEventListener("input", function() {
    getSearchSuggestions(this.value);
    currentFocus = -1;
  });
  // when ARROW DOWN/UP or ENTER key is pressed down
  inp.addEventListener("keydown", function(e) {
    if (e.keyCode == 40) {
      // ARROW DOWN
      currentFocus++;
      addClassActiveToAutoCompleteItem();
    } else if (e.keyCode == 38) {
      // ARROW UP
      currentFocus--;
      addClassActiveToAutoCompleteItem();
    } else if (e.keyCode == 13) {
      // ENTER
      if (currentFocus > -1) {
        e.preventDefault();
        clickOnAutoCompleteItem(currentFocus);
      }
    }
  });

  function addClassActiveToAutoCompleteItem() {
    let autoCompleteItems = getAutoCompleteItems();
    if (!autoCompleteItems) return false;
    removeClassActiveFromAutoCompleteItems(autoCompleteItems);
    if (currentFocus >= autoCompleteItems.length) currentFocus = 0;
    if (currentFocus < 0) currentFocus = (autoCompleteItems.length - 1);
    autoCompleteItems[currentFocus].classList.add("autocomplete-active");
  }

  function removeClassActiveFromAutoCompleteItems(autoCompleteItems) {
    for (let i = 0; i < autoCompleteItems.length; i++) {
      autoCompleteItems[i].classList.remove("autocomplete-active");
    }
  }

  function clickOnAutoCompleteItem() {
    // simulate a click on the "active" item
    let activeAutoCompleteItem = getActiveAutoCompleteItem();
    if (activeAutoCompleteItem) {
      activeAutoCompleteItem.getElementsByTagName("a")[0].click();
    }
  }

  document.addEventListener("click", removeAutoCompleteItems);
}

function init() {
  let inp = document.getElementById("query");
  initAutoComplete(inp);
}

init();
