var Boggle = { };

Boggle.BogglePage = function(pageElem)
{
   this.kDefaultSize = 5;

   this.fWidth = 0;
   this.fHeight = 0;

   this.fBoardWidthField = $("#boardWidth", pageElem);
   this.fBoardHeightField = $("#boardHeight", pageElem);
   this.fSolveLink = $("#solveLink", pageElem);
   this.fRandomizeLink = $("#randomizeLink", pageElem);
   this.fRandomizeEmptiesLink = $("#randomizeEmptiesLink", pageElem);
   this.fClearLink = $("#clearLink", pageElem);
   this.fBoard = $("#board", pageElem);
   this.fTableBody = $("tbody", this.fBoard);
   this.fWordList = $("#words", pageElem);

   this.fBoardWidthField.on("change", $.proxy(this.BoardWidthChanged, this));
   this.fBoardHeightField.on("change", $.proxy(this.BoardHeightChanged, this));
   this.fSolveLink.on("click", $.proxy(this.Solve, this));
   this.fRandomizeLink.on("click", $.proxy(this.RandomizeBoard, this));
   this.fRandomizeEmptiesLink.on("click",
    $.proxy(this.RandomizeEmpties, this));
   this.fClearLink.on("click", $.proxy(this.ClearBoard, this));

   this.fBoardWidthField.val(this.kDefaultSize);
   this.fBoardHeightField.val(this.kDefaultSize);

   this.ResizeBoard(this.kDefaultSize, this.kDefaultSize);
   this.RandomizeBoard();

   var bogglePage = this;

   this.fBoard.on("blur", "input", function()
   {
      // Pass the changed cell to the CellChanged handler.
      bogglePage.CellChanged.call(bogglePage, $(this));
   });

   this.fBoard.on("click", "td", function()
   {
      // Convert to input and select contents upon click.
      var cell = $(this);

      if (cell.children("input").size() > 0)
      {
         return;
      }

      var input = $('<input type="text" maxlength="1" />').val(cell.html());
      cell.empty().append(input);
      input.select();
   });
};

Boggle.BogglePage.prototype.BoardWidthChanged = function()
{
   this.ResizeBoard(parseInt(this.fBoardWidthField.val(), 10), this.fHeight);
};

Boggle.BogglePage.prototype.BoardHeightChanged = function()
{
   this.ResizeBoard(this.fWidth, parseInt(this.fBoardHeightField.val(), 10));
};

Boggle.BogglePage.prototype.CellChanged = function(cellInput)
{
   this.ClearSolution();

   // Force lower case and clear if invalid.

   var char = cellInput.val().toLowerCase();

   if (char < 'a' || char > 'z')
   {
      char = '';
   }

   // Move new value into parent <td />.

   cellInput.parent().empty().html(char);
};

Boggle.BogglePage.prototype.ClearBoard = function()
{
   this.ClearSolution();
   $("td", this.fBoard).html('');
};

Boggle.BogglePage.prototype.RandomizeEmpties = function()
{
   this.RandomizeBoard(false);
};

Boggle.BogglePage.prototype.RandomizeBoard = function(destructive)
{
   this.ClearSolution();

   if (_.isUndefined(destructive))
   {
      destructive = true;
   }

   var a = 'a'.charCodeAt(0);
   var cells;

   if (destructive)
   {
      cells = $("td", this.fBoard);
   }
   else
   {
      // Skip over cells that are already filled.
      cells = $("td", this.fBoard).filter(
       function() { return '' === $(this).html(); });
   }

   _.each(cells, function(cell)
   {
      $(cell).html(String.fromCharCode(a + Math.floor(Math.random() * 26)));
   });
};

Boggle.BogglePage.prototype.ResizeBoard = function(width, height)
{
   this.ClearSolution();

   var oldWidth = this.fWidth;
   var oldHeight = this.fHeight;
   this.fWidth = width;
   this.fHeight = height;

   if (0 === width || 0 === height)
   {
      this.fTableBody.empty();
      this.fWordList.empty();
      return;
   }

   var widthChange = width - oldWidth;
   var heightChange = height - oldHeight;

   if (heightChange < 0)
   {
      // Remove rows.
      this.fTableBody.children("tr").slice(heightChange).remove();
   }
   else if (heightChange > 0)
   {
      // Add rows with the new amount of columns.
      var rowString = '<tr>' + _.str.repeat('<td/>', width) + '</tr>';
      this.fTableBody.append(_.str.repeat(rowString, heightChange));
   }

   if (0 === widthChange || 0 === oldHeight)
   {
      // We're good -- all required work was done above.
      return;
   }

   // Reconcile columns within each row. We only need to look at $oldHeight
   //  rows - any new ones already have the correct number of columns.

   _.each(this.fTableBody.children("tr").slice(0, oldHeight), function(row)
   {
      row = $(row);

      if (widthChange < 0)
      {
         // Remove columns.
         row.children("td").slice(widthChange).remove();
      }
      else if (widthChange > 0)
      {
         row.append(_.str.repeat('<td/>', widthChange))
      }
   });
};

// Returns NxM encoded matrix string of characters.
Boggle.BogglePage.prototype.GetBoardMatrix = function()
{
   var rows = [];

   _.each(this.fTableBody.children("tr"), function (tr)
   {
      var chars = '';
      
      _.each($(tr).children("td"), function (td)
      {
         var char = $(td).html();

         if ('' === char)
         {
            throw new Error('Incomplete board.');
         }

         chars += char;
      });

      rows.push(chars);
   });

   return rows.join('\n');
};

Boggle.BogglePage.prototype.ClearSolution = function()
{
   this.fWordList.addClass('stale');
};

Boggle.BogglePage.prototype.Solve = function()
{
   try
   {
      var matrix = this.GetBoardMatrix();
   }
   catch (e)
   {
      alert("Solving requires that all cells be filled.");
      return;
   }

   $.post('/api/boggle/words', { '': matrix },
    $.proxy(this.ShowWords, this), 'json');
};

Boggle.BogglePage.prototype.ShowWords = function(data, textStatus, jqXhr)
{
   this.fWordList.empty().removeClass('stale');

   if (0 === data.length)
   {
      $("<li>No words found!</li>").appendTo(this.fWordList);
      return;
   }

   var newHtml = '';

   for (var i = 0; i < data.length; ++i)
   {
      newHtml += "<li>" + data[i] + "</li>";
   }

   this.fWordList.append(newHtml);
};
