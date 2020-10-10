// +build !cgo

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package syntax

import (
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

func (se *SyntaxEdit) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_NOTIFY, win.WM_COMMAND:
		switch win.HIWORD(uint32(wParam)) {
		case win.EN_CHANGE:
			se.textChangedPublisher.Publish()
		}
		// This is a horrible trick from MFC where we reflect the event back to the child.
		se.SendMessage(msg+ /*win.WM_REFLECT*/ win.WM_USER+0x1c00, wParam, lParam)
	}
	return se.WidgetBase.WndProc(hwnd, msg, wParam, lParam)
}

func makeSyntaxEdit(parent walk.Container) (*SyntaxEdit, error) {
	se := &SyntaxEdit{}
	err := walk.InitWidget(
		se,
		parent,
		"RICHEDIT50W",
		win.WS_CHILD|win.ES_MULTILINE|win.WS_VISIBLE|win.WS_VSCROLL|win.WS_BORDER|win.WS_HSCROLL|win.WS_TABSTOP|win.ES_WANTRETURN| /*win.ES_NOOLEDRAGDROP*/ 0x8,
		0,
	)
	if err != nil {
		return nil, err
	}
	return se, nil
}

// TODO: Implement func (se *SyntaxEdit) ApplyDPI(dpi int)
