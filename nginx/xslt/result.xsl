<?xml version="1.0"?>
<xsl:transform
  xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
  xmlns:html="http://www.w3.org/1999/xhtml"
  xmlns="http://www.w3.org/1999/xhtml"
  version="1.0"
>

<xsl:output method="html" omit-xml-declaration="yes"/>

<xsl:template match="/">
  <xsl:text disable-output-escaping="yes">&lt;!DOCTYPE html&gt;</xsl:text>
  <xsl:text disable-output-escaping="yes">&lt;html&gt;</xsl:text>
  <xsl:text disable-output-escaping="yes">&lt;head&gt;&lt;title&gt;Myblog Search&lt;/title&gt;&lt;head&gt;</xsl:text>
  <xsl:text disable-output-escaping="yes">&lt;link rel="stylesheet" href="/mystyle.css"&gt;</xsl:text>
  <xsl:text disable-output-escaping="yes">&lt;meta http-equiv="Content-Type" content="text/html; charset=utf-8"/&gt;</xsl:text>
  <xsl:text disable-output-escaping="yes">&lt;body&gt;</xsl:text>
    <xsl:apply-templates/>
  <xsl:text disable-output-escaping="yes">&lt;/body&gt;</xsl:text>
  <xsl:text disable-output-escaping="yes">&lt;/html&gt;</xsl:text>
</xsl:template>

<xsl:template match="result">
  <xsl:text disable-output-escaping="yes">&lt;!--#include file="/searchbar.html" --&gt;</xsl:text>
  <xsl:text disable-output-escaping="yes">&lt;table&gt;</xsl:text>
    <xsl:apply-templates select="hit"/>
  <xsl:text disable-output-escaping="yes">&lt;/table&gt;</xsl:text>
</xsl:template>

<xsl:template match="hit">
  <xsl:text disable-output-escaping="yes">&lt;tr&gt;</xsl:text>
    <xsl:text disable-output-escaping="yes">&lt;td&gt;</xsl:text>
      <xsl:if test="field[@name='thumbnail']">
        <xsl:text disable-output-escaping="yes">&lt;image src=&quot;</xsl:text>
        <xsl:value-of select="field[@name='thumbnail']"/>
        <xsl:text disable-output-escaping="yes">&quot;&gt;</xsl:text>
        <xsl:text disable-output-escaping="yes">&lt;/image&gt;</xsl:text>
      </xsl:if>
    <xsl:text disable-output-escaping="yes">&lt;/td&gt;</xsl:text>
    <xsl:text disable-output-escaping="yes">&lt;td&gt;</xsl:text>
      <xsl:if test="field[@name='url']">
        <xsl:text disable-output-escaping="yes">&lt;a href=&quot;</xsl:text>
        <xsl:value-of select="field[@name='url']"/>
        <xsl:text disable-output-escaping="yes">&quot;&gt;</xsl:text>
      </xsl:if>
      <xsl:if test="field[@name='title']">
        <xsl:value-of select="field[@name='title']"/>
      </xsl:if>
      <xsl:if test="field[@name='url']">
        <xsl:text disable-output-escaping="yes">&lt;/a&gt;</xsl:text>
      </xsl:if>
    <xsl:text disable-output-escaping="yes">&lt;/td&gt;</xsl:text>
    <xsl:text disable-output-escaping="yes">&lt;td&gt;</xsl:text>
      <xsl:apply-templates select="field[@name='body']"/>
    <xsl:text disable-output-escaping="yes">&lt;/td&gt;</xsl:text>
  <xsl:text disable-output-escaping="yes">&lt;/tr&gt;</xsl:text>
</xsl:template>

<xsl:template match="field">
  <xsl:for-each select="./node()">
    <xsl:choose>
      <xsl:when test="name() = 'hi'">
        <xsl:text disable-output-escaping="yes">&lt;b&gt;</xsl:text>
        <xsl:value-of select="text()"/>
        <xsl:text disable-output-escaping="yes">&lt;/b&gt;</xsl:text>
      </xsl:when>
      <xsl:when test="name() = 'sep'">
        <xsl:text> ... </xsl:text>
      </xsl:when>
      <xsl:otherwise>
        <xsl:value-of select="."/>
      </xsl:otherwise>
    </xsl:choose>
  </xsl:for-each>
</xsl:template>

</xsl:transform>
