/**
 * Created by zouxuan on 4/26/17.
 */

(function (global) {
    'use strict'

    function createBlock(name, value, content) {
        var item = elem('div', { 'class': 'block', draggable: true, 'data-name': name }, [name]);
        if (value !== undefined && value !== null) {
            if (typeof (value) === Number) {
                item.appendChild(elem('input', { type: 'number', value: value }));
            } else {
                item.appendChild(elem('input', { value: value }));
            }
            if (Array.isArray(content)) {
                if (content.length > 0) {
                    item.appendChild(elem('div', { 'class': 'block-container' }, content.map(function (block) {
                        return createBlock.apply(null, [block.cmd, block.params[0], block.contents]);
                    })));
                }
            }
        } else if (typeof content === 'string') {
            item.appendChild(document.createTextNode(' ' + content));
        }
        return item;
    }

    function blockContents(block) {
        var container = block.querySelector('.block-container');
        return container ? [].slice.call(container.children) : null;
    }

    function blockValue(block) {
        var input = block.querySelector('input');
        return Number(input.value);
    }

    function blockUnits(block) {
        if (block.children.length > 1 && block.lastChild.nodeType === Node.TEXT_NODE && block.lastChild.textContent) {
            return block.lastChild.textContent.slice(1);
        }
    }

    function blockScript(block) {
        var cmd = {};
        cmd.cmd = block.dataset.name;
        if (cmd.cmd.indexOf('>') !== -1) {
            cmd.cmd = cmd.cmd.split(' ')[0] + 'L'
        }
        if (cmd.cmd.indexOf('<') !== -1) {
            cmd.cmd = cmd.cmd.split(' ')[0] + 'S'
        }
        cmd.params = [];
        cmd.contents = [];
        var value = blockValue(block);
        if (value !== null) {
            cmd.params.push(value);
        }
        var contents = blockContents(block);
        if (contents) {
            cmd.contents = contents.map(blockScript);
        }
        return cmd;
    }

    function runBlocks(block) {
        block.forEach(function (block) {
            trigger('run', block);
        });
    }

    global.Block = {
        create: createBlock,
        value: blockValue,
        contents: blockContents,
        script: blockScript,
        run: runBlocks,
        trigger: trigger
    };
})(window);



