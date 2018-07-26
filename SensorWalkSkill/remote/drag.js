/**
 * Created by zouxuan on 4/27/17.
 */

(function(global) {
  'use strict';

  var dragTarget = null;
  var dragType = null;
  var scriptBlocks = [];

  function handleTouchMove(event) {
    console.log('touchmove', event);
  }

  function handleTouchLeave(event) {
    console.log('touchLeave', event);
  }

  function dragStart(event) {
    console.log('dragStart', event);

    if (!matches(event.target, '.block')) {
      return;
    }
    if (matches(event.target, '.menu .block')) {
      dragType = 'menu';
    } else {
      dragType = 'script';
    }
    event.target.classList.add('dragging');
    dragTarget = event.target;
    scriptBlocks = [].slice.call(
      document.querySelectorAll('.script .block:not(.dragging)')
    );
    event.dataTransfer.setData('text/html', event.target.outerHTML);
    if (matches(event.target, '.menu .block')) {
      event.dataTransfer.effectAllowed = 'copy';
    } else {
      event.dataTransfer.effectAllowed = 'move';
    }
  }

  function dragOver(event) {
    console.log('dragOver', event);

    if (
      !matches(event.target, '.menu, .menu *, .script, .script *, .content')
    ) {
      return;
    }
    if (event.preventDefault) {
      event.preventDefault();
    }
    if (dragType === 'menu') {
      event.dataTransfer.dropEffect = 'copy';
    } else {
      event.dataTransfer.dropEffect = 'move';
    }
    return false;
  }

  function drop(event) {
    console.log('drop', event);

    if (!matches(event.target, '.menu, .menu *, .script, .script *')) {
      return;
    }
    var dropTarget = closest(
      event.target,
      '.script .block-container, .script .block, .menu, .script'
    );
    var dropType = 'script';
    if (matches(dropTarget, '.menu')) {
      dropType = 'menu';
    }
    if (event.stopPropagation) {
      event.stopPropagation();
    }
    if (dragType === 'script' && dropType === 'menu') {
      trigger('blockRemoved', dragTarget.parentElement, dragTarget);
      dragTarget.parentElement.removeChild(dragTarget);
    } else if (dragType === 'script' && dropType === 'script') {
      if (matches(dropTarget, '.block')) {
        dropTarget.parentElement.insertBefore(
          dragTarget,
          dropTarget.nextSibling
        );
      } else {
        dropTarget.insertBefore(dragTarget, dropTarget.firstChildElement);
      }
      trigger('blockMoved', dropTarget, dragTarget);
    } else if (dragType === 'menu' && dropType === 'script') {
      var newNode = dragTarget.cloneNode(true);
      newNode.classList.remove('dragging');
      if (matches(dropTarget, '.block')) {
        dropTarget.parentElement.insertBefore(newNode, dropTarget.nextSibling);
      } else {
        dropTarget.insertBefore(newNode, dropTarget.firstChildElement);
      }
      trigger('blockAdded', dropTarget, newNode);
    }
  }

  function dragEnter(event) {
    console.log('dragEnter', event);

    if (matches(event.target, '.menu, .script, .content')) {
      event.target.classList.add('over');
      if (event.preventDefault) {
        event.preventDefault();
      } // Necessary. Allows us to drop.
    } else {
      if (!matches(event.target, '.menu *, .script *')) {
        _findAndRemoveClass('over');
        event.target.classList.remove('over');
      }
    }
    return false;
  }

  function _findAndRemoveClass(klass) {
    var elem = document.querySelector('.' + klass);
    if (elem) {
      elem.classList.remove(klass);
    }
  }

  function dragEnd(event) {
    _findAndRemoveClass('dragging');
    _findAndRemoveClass('over');
    _findAndRemoveClass('next');
  }

  // document.addEventListener("touchstart", handleDrag, false);
  // document.addEventListener("touchend", handleEnd, false);
  // document.addEventListener("touchcancel", handleCancel, false);
  document.addEventListener('touchleave', handleTouchLeave, false);
  document.addEventListener('touchmove', handleTouchMove, false);

  document.addEventListener('dragstart', dragStart, false);
  document.addEventListener('dragenter', dragEnter, false);
  document.addEventListener('dragover', dragOver, false);
  document.addEventListener('drag', function() {}, false);
  document.addEventListener('drop', drop, false);
  document.addEventListener('dragend', dragEnd, false);
})(window);
