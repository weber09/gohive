package tcliservice

func (c *TBinaryColumn) GetInterfaceArray() []interface{} {
	vals := c.GetValues()
	interfaceColumn := make([]interface{}, len(vals))
	for i := range vals {
		interfaceColumn[i] = vals[i]
	}
	return interfaceColumn
}

func (c *TBoolColumn) GetInterfaceArray() []interface{} {
	vals := c.GetValues()
	interfaceColumn := make([]interface{}, len(vals))
	for i := range vals {
		interfaceColumn[i] = vals[i]
	}
	return interfaceColumn
}

func (c *TByteColumn) GetInterfaceArray() []interface{} {
	vals := c.GetValues()
	interfaceColumn := make([]interface{}, len(vals))
	for i := range vals {
		interfaceColumn[i] = vals[i]
	}
	return interfaceColumn
}

func (c *TDoubleColumn) GetInterfaceArray() []interface{} {
	vals := c.GetValues()
	interfaceColumn := make([]interface{}, len(vals))
	for i := range vals {
		interfaceColumn[i] = vals[i]
	}
	return interfaceColumn
}

func (c *TI16Column) GetInterfaceArray() []interface{} {
	vals := c.GetValues()
	interfaceColumn := make([]interface{}, len(vals))
	for i := range vals {
		interfaceColumn[i] = vals[i]
	}
	return interfaceColumn
}

func (c *TI32Column) GetInterfaceArray() []interface{} {
	vals := c.GetValues()
	interfaceColumn := make([]interface{}, len(vals))
	for i := range vals {
		interfaceColumn[i] = vals[i]
	}
	return interfaceColumn
}

func (c *TI64Column) GetInterfaceArray() []interface{} {
	vals := c.GetValues()
	interfaceColumn := make([]interface{}, len(vals))
	for i := range vals {
		interfaceColumn[i] = vals[i]
	}
	return interfaceColumn
}

func (c *TStringColumn) GetInterfaceArray() []interface{} {
	vals := c.GetValues()
	interfaceColumn := make([]interface{}, len(vals))
	for i := range vals {
		interfaceColumn[i] = vals[i]
	}
	return interfaceColumn
}
