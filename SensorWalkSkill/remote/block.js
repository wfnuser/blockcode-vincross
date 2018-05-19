/**
 * Created by zouxuan on 4/26/17.
 */

(function (global) {
    'use strict'

    var Conditions = ['isRed', 'isBright']

    function createBlock(name, value, content) {
        var item = elem('div', { 'class': 'block', draggable: true, 'data-name': name }, [name]);
        if (value !== undefined && value !== null) {
            if (typeof (value) === Number) {
                item.appendChild(elem('input', { type: 'number', value: value }));
            } else {
                item.appendChild(elem('input', { value: value }));
            }
            if (Array.isArray(content)) {
                item.appendChild(elem('div', { 'class': 'container' }, content.map(function (block) {
                    return createBlock.apply(null, block);
                })));
            }
        } else if (typeof content === 'string') {
            item.appendChild(document.createTextNode(' ' + content));
        }
        return item;
    }

    function blockContents(block) {
        var container = block.querySelector('.container');
        return container ? [].slice.call(container.children) : null;
    }

    function blockValue(block) {
        var input = block.querySelector('input');
        if (input && Conditions.indexOf(input.value) >= 0) {
            return input.value;
        } else {
            return Number(input.value);
        }
    }

    function blockUnits(block) {
        if (block.children.length > 1 && block.lastChild.nodeType === Node.TEXT_NODE && block.lastChild.textContent) {
            return block.lastChild.textContent.slice(1);
        }
    }

    function blockScript(block) {
        var script = [];
        var cmd = {};
        cmd.cmd = block.dataset.name;
        cmd.params = [];
        cmd.contents = [];
        var value = blockValue(block);
        if (value !== null) {
            if (cmd.cmd === 'If') {
                if (Conditions.indexOf(value) >= 0) {
                    cmd.params.push(Conditions.indexOf(value))
                }
            } else {
                cmd.params.push(value);
            }
        }
        var contents = blockContents(block);
        if (contents) {
            cmd.contents.push(contents.map(blockScript));
        }
        script.push(cmd)
        return script.filter(function (notNull) {
            return notNull !== null;
        });
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



