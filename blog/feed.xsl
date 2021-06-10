<?xml version="1.0"?>
<xsl:stylesheet version="1.0"
  xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
  xmlns:atom="http://www.w3.org/2005/Atom"
  xmlns:openSearch="http://a9.com/-/spec/opensearchrss/1.0/"
  xmlns:blogger="http://schemas.google.com/blogger/2008"
  xmlns:georss="http://www.georss.org/georss"
  xmlns:gd="http://schemas.google.com/g/2005"
  xmlns:thr="http://purl.org/syndication/thread/1.0"
  xmlns:media="http://search.yahoo.com/mrss/"
>

<xsl:output method="text" omit-xml-declaration="yes"/>

<xsl:param name="namespace" select="'myblog'"/>
<xsl:param name="document-type" select="'myblog'"/>

<xsl:template name="doublequotes">
  <xsl:param name="text" select="."/>
  <xsl:variable name="quot">"</xsl:variable>
  <xsl:choose>
    <xsl:when test="contains($text, $quot)">
      <xsl:value-of select="substring-before($text, $quot)"/>
      <xsl:text>&amp;quot;</xsl:text>
      <xsl:call-template name="doublequotes">
        <xsl:with-param name="text" select="substring-after($text, $quot)"/>
      </xsl:call-template>
    </xsl:when>
    <xsl:otherwise>
      <xsl:value-of select="$text"/>
    </xsl:otherwise>
  </xsl:choose>
</xsl:template>

<xsl:template name="htmltags">
  <xsl:param name="text" select="."/>
  <xsl:choose>
    <xsl:when test="contains($text, '&lt;') and contains(substring-after($text, '&lt;'), '&gt;')">
      <xsl:value-of select="substring-before($text, '&lt;')"/>
      <xsl:call-template name="htmltags">
        <xsl:with-param name="text" select="substring-after(substring-after($text, '&lt;'), '&gt;')"/>
      </xsl:call-template>
    </xsl:when>
    <xsl:otherwise>
      <xsl:value-of select="$text"/>
    </xsl:otherwise>
  </xsl:choose>
</xsl:template>

<xsl:template match="/">
  <xsl:apply-templates select="atom:feed"/>
</xsl:template>

<xsl:template match="atom:feed">
  <xsl:text>[</xsl:text>
  <xsl:apply-templates select="atom:entry"/>
  <xsl:text>]&#xA;</xsl:text>
</xsl:template>

<xsl:template match="atom:entry">
  <xsl:variable name="id" select="substring-after(atom:id, 'post-')"/>
    <xsl:text>{</xsl:text>
      <xsl:text>"put":"id:</xsl:text>
      <xsl:value-of select="$namespace"/>
      <xsl:text>:</xsl:text>
      <xsl:value-of select="$document-type"/>
      <xsl:text>::</xsl:text>
      <xsl:value-of select="$id"/>",
      <xsl:text>"fields":{</xsl:text>
        <xsl:text>"language":"</xsl:text>
        <xsl:if test="atom:category[@term='Australia']">
          <xsl:text>en</xsl:text>
        </xsl:if>
        <xsl:if test="not(atom:category[@term='Australia'])">
          <xsl:text>zh-TW</xsl:text>
        </xsl:if>
        <xsl:text>",</xsl:text>
        <xsl:text>"id":"</xsl:text>
        <xsl:value-of select="$id"/>
        <xsl:text>",</xsl:text>
        <xsl:apply-templates select="atom:link[@rel='alternate']|atom:title|atom:content|media:thumbnail"/>
      <xsl:text>}</xsl:text>
    <xsl:text>}</xsl:text>
  <xsl:if test="position() != last()"><xsl:text>,</xsl:text></xsl:if>
</xsl:template>

<xsl:template match="atom:link">
  <xsl:variable name="url" select="@href"/>
  <xsl:text>"url":"</xsl:text>
  <xsl:value-of select="$url"/>
  <xsl:text>"</xsl:text>
  <xsl:if test="position() != last()"><xsl:text>,</xsl:text></xsl:if>
</xsl:template>

<xsl:template match="atom:title">
  <xsl:text>"title":"</xsl:text>
  <xsl:value-of select="."/>
  <xsl:text>"</xsl:text>
  <xsl:if test="position() != last()"><xsl:text>,</xsl:text></xsl:if>
</xsl:template>

<xsl:template match="atom:content">
  <xsl:variable name="escaped">
    <xsl:call-template name="doublequotes"/>
  </xsl:variable>
  <xsl:text>"body":"</xsl:text>
  <xsl:call-template name="htmltags">
    <xsl:with-param name="text" select="$escaped"/>
  </xsl:call-template>
  <xsl:text>"</xsl:text>
  <xsl:if test="position() != last()"><xsl:text>,</xsl:text></xsl:if>
</xsl:template>

<xsl:template match="media:thumbnail">
  <xsl:text>"thumbnail":"</xsl:text>
  <xsl:value-of select="@url"/>
  <xsl:text>"</xsl:text>
  <xsl:if test="position() != last()"><xsl:text>,</xsl:text></xsl:if>
</xsl:template>

</xsl:stylesheet>
